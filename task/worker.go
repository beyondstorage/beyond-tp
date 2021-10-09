package task

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"

	"github.com/beyondstorage/beyond-tp/proto"
)

type Worker struct {
	c *Client

	src types.Storager
	dst types.Storager

	wg *sync.WaitGroup
}

func HandleJob(c *Client, j *Job) {
	// TODO: we need to set src and dst here.
	w := &Worker{
		c: c,
	}

	ctx := context.Background()
	err := w.Handle(ctx, j)
	if err != nil {
		c.errch <- err
	}
	_, err = w.c.gc.FinishJob(ctx, &proto.FinishJobRequest{
		TaskId: j.TaskId,
		JobId:  j.Id,
		Status: 0,
	})
	if err != nil {
		c.errch <- err
	}
}

func (w *Worker) Sync(ctx context.Context, j *Job) (err error) {
	_, err = w.c.gc.CreateJob(ctx, ToProtoJob(j))
	if err != nil {
		return
	}
	_, err = w.c.gc.WaitJob(ctx, &proto.WaitJobRequest{
		TaskId: j.TaskId,
		JobId:  j.Id,
	})
	if err != nil {
		return
	}
	return
}

func (w *Worker) Async(ctx context.Context, j *Job) (err error) {
	w.wg.Add(1)
	_, err = w.c.gc.CreateJob(ctx, ToProtoJob(j))
	if err != nil {
		return
	}

	// TODO: Do we need to use a pool to wait job?
	go func() {
		defer w.wg.Done()

		_, err = w.c.gc.WaitJob(ctx, &proto.WaitJobRequest{
			TaskId: j.TaskId,
			JobId:  j.Id,
		})
		if err != nil {
			w.c.errch <- err
		}
	}()
	return
}

func (w *Worker) Await() {
	w.wg.Wait()
}

func (w *Worker) Handle(ctx context.Context, j *Job) (err error) {
	var h func(ctx context.Context, j *Job) error

	switch j.Type {
	case JobTypeCopyDir:
		h = w.HandleCopyDir
	case JobTypeCopySmallFile:
		h = w.HandleCopySmallFile
	case JobTypeCopyLargeFile:
		h = w.HandleCopyLargeFile
	case JobTypeCopyPart:
		h = w.HandleCopyPart
	}

	return h(ctx, j)
}

func (w *Worker) HandleCopyDir(ctx context.Context, j *Job) (err error) {
	ij := &CopyDirJob{}
	MustUnmarshal(j.Content, ij)

	it, err := w.src.List(ij.SrcPath)
	if err != nil {
		return fmt.Errorf("list source: %w", err)
	}

	for {
		o, err := it.Next()
		if err != nil && errors.Is(err, types.IterateDone) {
			break
		}
		if err != nil {
			return fmt.Errorf("list next: %w", err)
		}
		if o.Mode.IsDir() {
			cj := CopyDirJob{
				SrcPath: o.Path,
				DstPath: o.Path,
			}

			err = w.Async(ctx, NewJob(j.TaskId, JobTypeCopyDir, MustMarshal(cj)))
			if err != nil {
				return err
			}
			continue
		}
		// TODO: we need to support stream file either.
		size := o.MustGetContentLength()
		if size > LargeFileBoundary {
			cj := CopyLargeFileJob{
				SrcPath: o.Path,
				DstPath: o.Path,
				Size:    size,
			}

			err = w.Async(ctx, NewJob(j.TaskId, JobTypeCopyLargeFile, MustMarshal(cj)))
			if err != nil {
				return err
			}
			continue
		}

		cj := CopySmallFileJob{
			SrcPath: o.Path,
			DstPath: o.Path,
			Size:    size,
		}
		err = w.Async(ctx, NewJob(j.TaskId, JobTypeCopySmallFile, MustMarshal(cj)))
		if err != nil {
			return err
		}
		continue
	}

	w.Await()
	return nil
}

func (w *Worker) HandleCopySmallFile(ctx context.Context, j *Job) (err error) {
	ij := &CopySmallFileJob{}
	MustUnmarshal(j.Content, ij)

	pr, pw := io.Pipe()

	go func() {
		_, err := w.src.Read(ij.SrcPath, pw, pairs.WithSize(ij.Size))
		if err != nil {
			return
		}
	}()

	_, err = w.dst.Write(ij.DstPath, pr, ij.Size)
	if err != nil {
		return
	}
	return
}

func (w *Worker) HandleCopyLargeFile(ctx context.Context, j *Job) (err error) {
	ij := &CopyLargeFileJob{}
	MustUnmarshal(j.Content, ij)

	mu := w.dst.(types.Multiparter)

	o, err := mu.CreateMultipart(ij.DstPath)
	if err != nil {
		return
	}

	index := 0
	offset := int64(0)
	totalSize := ij.Size
	partSize := int64(64 * 1024 * 1024) // Use 64 MiB as part size.
	parts := make([]*types.Part, 0)
	partJobs := make(map[int]string) // index -> job id.

	for offset < totalSize {
		size := partSize
		if offset+size > totalSize {
			size = totalSize - offset
		}

		cp := &CopyPartJob{
			SrcPath:     ij.SrcPath,
			DstPath:     ij.DstPath,
			MultipartId: o.MustGetMultipartID(),
			Size:        size,
			Index:       index,
			Offset:      offset,
		}
		partJob := NewJob(j.TaskId, JobTypeCopyPart, MustMarshal(cp))
		err = w.Async(ctx, partJob)
		if err != nil {
			return err
		}

		parts = append(parts, &types.Part{
			Index: index,
			Size:  size,
		})
		partJobs[index] = partJob.Id
		offset += size
		index += 1
	}

	// Wait for all async job finished.
	w.Await()

	for i := 0; i < index; i++ {
		entry, err := w.c.gc.GetMeta(ctx, &proto.MetaKey{
			TaskId: j.TaskId,
			JobId:  partJobs[i],
			Key:    "etag",
		})
		if err != nil {
			return err
		}
		parts[i].ETag = entry.Value
	}

	err = mu.CompleteMultipart(o, parts)
	if err != nil {
		return
	}
	return
}

func (w *Worker) HandleCopyPart(ctx context.Context, j *Job) (err error) {
	ij := &CopyPartJob{}
	MustUnmarshal(j.Content, ij)

	mu := w.dst.(types.Multiparter)

	pr, pw := io.Pipe()

	go func() {
		_, err := w.src.Read(ij.SrcPath, pw, pairs.WithSize(ij.Size), pairs.WithOffset(ij.Offset))
		if err != nil {
			return
		}
	}()

	o := w.dst.Create(ij.DstPath, pairs.WithMultipartID(ij.MultipartId))

	_, part, err := mu.WriteMultipart(o, pr, ij.Size, ij.Index)
	if err != nil {
		return
	}

	_, err = w.c.gc.SetMeta(ctx, &proto.MetaEntry{
		TaskId: j.Id,
		JobId:  j.Id,
		Key:    "etag",
		Value:  part.ETag,
	})
	if err != nil {
		return
	}
	return nil
}
