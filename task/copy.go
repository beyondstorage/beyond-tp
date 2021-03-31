package task

import (
	"context"
	"fmt"

	ps "github.com/aos-dev/go-storage/v3/pairs"
	"github.com/aos-dev/go-storage/v3/pkg/iowrap"
	"github.com/aos-dev/go-storage/v3/types"
	protobuf "github.com/golang/protobuf/proto"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/proto"
)

const (
	defaultMultipartThreshold int64 = 1024 * 1024 * 1024 // 1G
	defaultMultipartPartSize  int64 = 128 * 1024 * 1024  // 128M
)

func (rn *Runner) HandleCopyDir(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger
	arg := msg.(*proto.CopyDir)

	store := rn.storages[arg.Src]

	it, err := store.List(arg.SrcPath)
	if err != nil {
		logger.Error("storage list failed",
			zap.Error(err),
			zap.String("store", store.String()),
			zap.String("path", arg.SrcPath))
		return err
	}

	for {
		o, err := it.Next()
		if err == types.IterateDone {
			break
		}
		if err != nil {
			logger.Error("get next object failed",
				zap.Error(err),
				zap.String("store", store.String()))
			return err
		}

		// if obj is dir and not recursive, skip directly
		if o.GetMode().IsDir() && !arg.Recursive {
			continue
		}

		job := proto.NewJob()
		// set job attr separately for dir and file
		if o.GetMode().IsDir() {
			content, err := protobuf.Marshal(&proto.CopyDir{
				Src:       arg.Src,
				Dst:       arg.Dst,
				SrcPath:   o.Path,
				DstPath:   o.Path,
				Recursive: true,
			})
			if err != nil {
				panic("marshal failed")
			}
			job.Type = TypeCopyDir
			job.Content = content
		} else {
			content, err := protobuf.Marshal(&proto.CopyFile{
				Src:     arg.Src,
				Dst:     arg.Dst,
				SrcPath: o.Path,
				DstPath: o.Path,
			})
			if err != nil {
				panic("marshal failed")
			}
			job.Type = TypeCopyFile
			job.Content = content
		}

		err = rn.Async(ctx, &job)
		if err != nil {
			logger.Error("async copy dir job failed",
				zap.Error(err),
				zap.String("object ID", o.ID),
				zap.String("store", store.String()))
			return err
		}
	}

	if err = rn.Await(ctx); err != nil {
		logger.Error("await copy dir job failed",
			zap.Error(err),
			zap.String("parent job", rn.j.Id),
			zap.String("store", store.String()))
		return err
	}

	logger.Info("copy dir job finished.")
	return nil
}

func (rn *Runner) HandleCopyFile(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger
	arg := msg.(*proto.CopyFile)

	src := rn.storages[arg.Src]
	dst := rn.storages[arg.Dst]

	obj, err := src.Stat(arg.SrcPath)
	if err != nil {
		return err
	}
	size, ok := obj.GetContentLength()
	if !ok {
		return fmt.Errorf("object %s size not set", arg.SrcPath)
	}

	job := proto.NewJob()
	if _, ok := dst.(types.Multiparter); ok && size > defaultMultipartThreshold {
		content, err := protobuf.Marshal(&proto.CopyMultipartFile{
			Src:     arg.Src,
			Dst:     arg.Dst,
			SrcPath: arg.SrcPath,
			DstPath: arg.DstPath,
			Size:    size,
		})
		if err != nil {
			panic("marshal failed")
		}

		job.Type = TypeCopyMultipartFile
		job.Content = content
	} else {
		content, err := protobuf.Marshal(&proto.CopySingleFile{
			Src:     arg.Src,
			Dst:     arg.Dst,
			SrcPath: arg.SrcPath,
			DstPath: arg.DstPath,
			Size:    size,
		})
		if err != nil {
			panic("marshal failed")
		}

		job.Type = TypeCopySingleFile
		job.Content = content
	}

	logger.Info("copy file",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))

	if err := rn.Sync(ctx, &job); err != nil {
		logger.Error("sync copy file job failed",
			zap.Error(err),
			zap.String("parent job", rn.j.Id),
			zap.String("src", src.String()),
			zap.String("dst", dst.String()))
		return err
	}

	return nil
}

func (rn *Runner) HandleCopySingleFile(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger

	arg := msg.(*proto.CopySingleFile)

	src := rn.storages[arg.Src]
	dst := rn.storages[arg.Dst]

	r, w := iowrap.Pipe()

	go func() {
		_, err := src.Read(arg.SrcPath, w)
		if err != nil {
			logger.Error("src read failed", zap.Error(err))
		}
	}()

	_, err := dst.Write(arg.DstPath, r, arg.Size)
	if err != nil {
		logger.Error("write single file failed", zap.Error(err))
		return err
	}

	defer func() {
		err = r.Close()
		if err != nil {
			return
		}
	}()

	logger.Info("copy single file",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))
	return nil
}
func (rn *Runner) HandleCopyMultipartFile(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger

	arg := msg.(*proto.CopyMultipartFile)

	dst := rn.storages[arg.Dst]
	multiparter := dst.(types.Multiparter)
	obj, err := multiparter.CreateMultipartWithContext(ctx, arg.DstPath)
	if err != nil {
		logger.Error("create multipart failed", zap.Error(err), zap.String("dst", dst.String()))
		return err
	}

	partSize, err := calculatePartSize(obj, arg.Size)
	if err != nil {
		logger.Error("calculate part size failed", zap.Error(err))
		return err
	}

	var offset int64
	var index uint32
	parts := make([]*types.Part, 0)
	for {
		// handle size for the last part
		if offset+partSize > arg.Size {
			partSize = arg.Size - offset
		}

		job := proto.NewJob()
		content, err := protobuf.Marshal(&proto.CopyMultipart{
			Src:         arg.Src,
			Dst:         arg.Dst,
			SrcPath:     arg.SrcPath,
			DstPath:     arg.DstPath,
			Size:        partSize,
			Index:       index,
			Offset:      offset,
			MultipartId: obj.MustGetMultipartID(),
		})
		if err != nil {
			panic("marshal failed")
		}

		job.Type = TypeCopyMultipart
		job.Content = content

		if err = rn.Async(ctx, &job); err != nil {
			logger.Error("async copy multipart failed", zap.Error(err), zap.String("parent job", rn.j.Id))
			return err
		}

		parts = append(parts, &types.Part{
			Index: int(index),
			Size:  partSize,
		})

		offset += partSize
		if offset >= arg.Size {
			break
		}
		index++
	}

	if err := rn.Await(ctx); err != nil {
		logger.Error("await copy multipart file failed", zap.Error(err), zap.String("parent job", rn.j.Id))
		return err
	}

	if err = multiparter.CompleteMultipartWithContext(ctx, obj, parts); err != nil {
		return err
	}

	// Send task and wait for response.
	logger.Info("copy multipart",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))
	return nil
}

func (rn *Runner) HandleCopyMultipart(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger

	arg := msg.(*proto.CopyMultipart)

	src := rn.storages[arg.Src]
	dst := rn.storages[arg.Dst]
	multipart, ok := dst.(types.Multiparter)
	if !ok {
		logger.Warn("storage does not implement Multiparter",
			zap.String("storage", dst.String()))
		return fmt.Errorf("not supported")
	}

	r, w := iowrap.Pipe()

	go func() {
		_, err := src.Read(arg.SrcPath, w, ps.WithSize(arg.Size), ps.WithOffset(arg.Offset))
		if err != nil {
			logger.Error("src read failed", zap.Error(err), zap.String("src", arg.SrcPath))
		}
	}()

	o := dst.Create(arg.DstPath, ps.WithMultipartID(arg.MultipartId))
	_, err := multipart.WriteMultipart(o, r, arg.Size, int(arg.Index))
	if err != nil {
		logger.Error("write multipart failed", zap.Error(err), zap.String("dst", arg.DstPath))
		return err
	}

	defer func() {
		err = r.Close()
		if err != nil {
			return
		}
	}()

	logger.Info("copy multipart",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))
	return nil
}
