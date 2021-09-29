package task

import (
	"context"
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

func StartWorker(c *Client, j *Job) {
	w := &Worker{}

	ctx := context.Background()
	err := w.Handle(ctx, j)
	if err != nil {
		c.errch <- err
	}
}

func (w *Worker) Sync()  {}
func (w *Worker) Async() {}
func (w *Worker) Await() {}

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
	return
}
func (w *Worker) HandleCopySmallFile(ctx context.Context, j *Job) (err error) {
	return
}
func (w *Worker) HandleCopyLargeFile(ctx context.Context, j *Job) (err error) {
	return
}
func (w *Worker) HandleCopyPart(ctx context.Context, j *Job) (err error) {
	ij := ParseCopyPartJob(j.Content)

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
		Key:   "etag",
		Value: part.ETag,
	})
	if err != nil {
		return
	}
	return nil
}
