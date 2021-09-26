package task

import (
	"context"
	"fmt"
	"io"

	ps "github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
	protobuf "github.com/golang/protobuf/proto"
	"go.uber.org/zap"

	"github.com/beyondstorage/beyond-tp/models"
)

const (
	defaultMultipartThreshold int64 = 1024 * 1024 * 1024 // 1G
	defaultMultipartPartSize  int64 = 128 * 1024 * 1024  // 128M
)

func (rn *runner) HandleCopyDir(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger
	arg := msg.(*models.CopyDirJob)

	store := rn.storages[arg.Src]

	it, err := store.List(arg.SrcPath, ps.WithListMode(types.ListModeDir))
	if err != nil {
		logger.Error("storage list",
			zap.String("store", store.String()),
			zap.String("path", arg.SrcPath),
			zap.Error(err))
		return err
	}

	for {
		o, err := it.Next()
		if err == types.IterateDone {
			break
		}
		if err != nil {
			logger.Error("get next object",
				zap.String("store", store.String()),
				zap.Error(err))
			return err
		}

		// if obj is dir and not recursive, skip directly
		if o.GetMode().IsDir() && !arg.Recursive {
			continue
		}

		var job *models.Job
		// set job attr separately for dir and file
		if o.GetMode().IsDir() {
			job = models.NewJob(models.JobType_CopyDir, &models.CopyDirJob{
				Src:       arg.Src,
				Dst:       arg.Dst,
				SrcPath:   o.Path,
				DstPath:   o.Path,
				Recursive: true,
			})
		} else {
			job = models.NewJob(models.JobType_CopyFile, &models.CopyFileJob{
				Src:     arg.Src,
				Dst:     arg.Dst,
				SrcPath: o.Path,
				DstPath: o.Path,
			})
		}

		err = rn.Async(ctx, job)
		if err != nil {
			logger.Error("async new job",
				zap.String("object_id", o.ID),
				zap.String("store", store.String()),
				zap.Error(err))
			return err
		}
	}

	if err = rn.Await(ctx); err != nil {
		logger.Error("await copy dir job",
			zap.String("parent job", rn.j.Id),
			zap.String("store", store.String()),
			zap.Error(err))
		return err
	}

	logger.Info("copy dir job finished.")
	return nil
}

func (rn *runner) HandleCopyFile(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger
	arg := msg.(*models.CopyFileJob)

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

	var job *models.Job
	if _, ok := dst.(types.Multiparter); ok && size > defaultMultipartThreshold {
		job = models.NewJob(models.JobType_CopyMultipartFile, &models.CopyMultipartFileJob{
			Src:     arg.Src,
			Dst:     arg.Dst,
			SrcPath: arg.SrcPath,
			DstPath: arg.DstPath,
			Size:    size,
		})
	} else {
		job = models.NewJob(models.JobType_CopySingleFile, &models.CopySingleFileJob{
			Src:     arg.Src,
			Dst:     arg.Dst,
			SrcPath: arg.SrcPath,
			DstPath: arg.DstPath,
			Size:    size,
		})
	}

	logger.Info("copy file",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))

	if err := rn.Sync(ctx, job); err != nil {
		logger.Error("copy file job",
			zap.Error(err),
			zap.String("parent job", rn.j.Id),
			zap.String("src", src.String()),
			zap.String("dst", dst.String()))
		return err
	}

	return nil
}

func (rn *runner) HandleCopySingleFile(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger

	arg := msg.(*models.CopySingleFileJob)

	logger.Debug("start copy single file", zap.String("src", arg.SrcPath), zap.String("dst", arg.DstPath))
	src := rn.storages[arg.Src]
	dst := rn.storages[arg.Dst]

	r, w := io.Pipe()

	go func() {
		defer func() {
			err := w.Close()
			if err != nil {
				logger.Error("close pipe writer", zap.Error(err))
			}
		}()
		_, err := src.Read(arg.SrcPath, w)
		if err != nil {
			logger.Error("src read failed", zap.Error(err))
		}
	}()

	logger.Debug("start write dst", zap.String("file", arg.DstPath))
	_, err := dst.Write(arg.DstPath, r, arg.Size)
	if err != nil {
		logger.Error("write single file failed", zap.Error(err))
		return err
	}

	logger.Info("copy single file",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))
	return nil
}

func (rn *runner) HandleCopyMultipartFile(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger

	arg := msg.(*models.CopyMultipartFileJob)

	dst := rn.storages[arg.Dst]
	multiparter := dst.(types.Multiparter)
	obj, err := multiparter.CreateMultipartWithContext(ctx, arg.DstPath)
	if err != nil {
		logger.Error("create multipart",
			zap.String("dst", dst.String()), zap.Error(err))
		return err
	}

	partSize, err := calculatePartSize(dst.Metadata(), arg.Size)
	if err != nil {
		logger.Error("calculate part size", zap.Error(err))
		return err
	}

	var offset int64
	var index uint32
	parts := make([]*types.Part, 0)
	indexToJob := make(map[uint32]string, 0)
	for {
		// handle size for the last part
		if offset+partSize > arg.Size {
			partSize = arg.Size - offset
		}

		job := models.NewJob(models.JobType_CopyMultipart, &models.CopyMultipartJob{
			Src:         arg.Src,
			Dst:         arg.Dst,
			SrcPath:     arg.SrcPath,
			DstPath:     arg.DstPath,
			Size:        partSize,
			Index:       index,
			Offset:      offset,
			MultipartId: obj.MustGetMultipartID(),
		})
		// register job ID with index
		indexToJob[index] = job.Id

		if err = rn.Async(ctx, job); err != nil {
			logger.Error("async copy multipart",
				zap.String("parent job", rn.j.Id), zap.Error(err))
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
		logger.Error("await copy multipart file",
			zap.String("parent job", rn.j.Id), zap.Error(err))
		return err
	}

	// aggregation part from metadata after await (to ensure metadata available)
	for i, jobID := range indexToJob {
		result, err := rn.grpcClient.GetJobMetadata(ctx, &models.GetJobMetadataRequest{
			JobId: jobID,
		})
		if err != nil {
			logger.Error("get copy multipart job metadata", zap.String("job", jobID), zap.Error(err))
			return err
		}

		metadata := models.WriteMultipartJobMetadata{}
		_ = protobuf.Unmarshal(result.Metadata, &metadata)
		parts[int(i)].ETag = metadata.Etag

		// delete job metadata after used
		_, err = rn.grpcClient.DeleteJobMetadata(ctx, &models.DeleteJobMetadataRequest{
			JobId: jobID,
		})
		if err != nil {
			logger.Warn("delete copy multipart job metadata failed", zap.String("job", jobID), zap.Error(err))
		}
	}

	if err = multiparter.CompleteMultipartWithContext(ctx, obj, parts); err != nil {
		return err
	}

	// Send task and wait for response.
	logger.Info("copy multipart file",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))
	return nil
}

func (rn *runner) HandleCopyMultipart(ctx context.Context, msg protobuf.Message) error {
	logger := rn.logger

	arg := msg.(*models.CopyMultipartJob)

	src := rn.storages[arg.Src]
	dst := rn.storages[arg.Dst]
	multipart, ok := dst.(types.Multiparter)
	if !ok {
		logger.Warn("storage does not implement Multiparter",
			zap.String("storage", dst.String()))
		return fmt.Errorf("not supported")
	}

	r, w := io.Pipe()

	go func() {
		defer func() {
			err := w.Close()
			if err != nil {
				logger.Error("close pipe writer", zap.Error(err))
			}
		}()
		_, err := src.Read(arg.SrcPath, w, ps.WithSize(arg.Size), ps.WithOffset(arg.Offset))
		if err != nil {
			logger.Error("src read",
				zap.String("src", arg.SrcPath), zap.Error(err))
		}
	}()

	o := dst.Create(arg.DstPath, ps.WithMultipartID(arg.MultipartId))
	_, part, err := multipart.WriteMultipart(o, r, arg.Size, int(arg.Index))
	if err != nil {
		logger.Error("write multipart",
			zap.String("dst", arg.DstPath), zap.Error(err))
		return err
	}

	result, _ := protobuf.Marshal(&models.WriteMultipartJobMetadata{
		Etag: part.ETag,
	})
	_, err = rn.grpcClient.SetJobMetadata(ctx, &models.SetJobMetadataRequest{
		JobId:    rn.j.Id,
		Metadata: result,
	})

	if err != nil {
		return err
	}
	logger.Info("copy multipart",
		zap.String("from", arg.SrcPath),
		zap.String("to", arg.DstPath))
	return nil
}
