package task

import (
	"context"
	"fmt"
	"sync"

	"github.com/beyondstorage/go-storage/v4/types"
	protobuf "github.com/golang/protobuf/proto"
	"go.uber.org/zap"

	"github.com/beyondstorage/beyond-tp/models"
)

type runner struct {
	j *models.Job

	grpcClient models.AgentClient
	storages   []types.Storager
	wg         *sync.WaitGroup

	logger *zap.Logger
}

func newRunner(j *models.Job) (*runner, error) {
	rn := &runner{
		j:     j,
		agent: a,

		grpcClient: a.grpcClient,
		storages:   a.storages,
		wg:         &sync.WaitGroup{},
		logger:     a.logger,
	}

	return rn, nil
}

func (rn *runner) Handle() {
	rn.logger.Info("runner start job",
		zap.String("id", rn.j.Id),
		zap.String("type", rn.j.Type.String()))

	ctx := context.Background()

	var fn func(ctx context.Context, msg protobuf.Message) error
	var t protobuf.Message

	switch rn.j.Type {
	case models.JobType_CopyDir:
		t = &models.CopyDirJob{}
		fn = rn.HandleCopyDir
	case models.JobType_CopyFile:
		t = &models.CopyFileJob{}
		fn = rn.HandleCopyFile
	case models.JobType_CopySingleFile:
		t = &models.CopySingleFileJob{}
		fn = rn.HandleCopySingleFile
	case models.JobType_CopyMultipartFile:
		t = &models.CopyMultipartFileJob{}
		fn = rn.HandleCopyMultipartFile
	case models.JobType_CopyMultipart:
		t = &models.CopyMultipartJob{}
		fn = rn.HandleCopyMultipart
	default:
		panic("not support job type")
	}

	err := protobuf.Unmarshal(rn.j.Content, t)
	if err != nil {
		panic(fmt.Errorf("job unmarshal, %w", err))
	}

	err = fn(ctx, t)
	if err != nil {
		rn.logger.Error("handle job", zap.Error(err), zap.String("type", rn.j.Type.String()))
	}

	// Send JobReply after the job has been handled.
	err = rn.Finish(ctx, err)
	if err != nil {
		rn.logger.Error("runner finish", zap.Error(err), zap.String("type", rn.j.Type.String()))
	}
}

func (rn *runner) Async(ctx context.Context, job *models.Job) (err error) {
	logger := rn.logger

	rn.wg.Add(1)

	_, err = rn.grpcClient.CreateJob(ctx, &models.CreateJobRequest{Job: job})
	if err != nil {
		return
	}
	go func() {
		defer rn.wg.Done()

		_, err := rn.grpcClient.WaitJob(ctx, &models.WaitJobRequest{JobId: job.Id})
		if err != nil {
			logger.Error("wait job", zap.Error(err))
		}
	}()

	logger.Info("runner publish async job",
		zap.String("job", job.Id))
	return
}

func (rn *runner) Await(ctx context.Context) (err error) {
	rn.wg.Wait()

	rn.logger.Info("runner finish await job",
		zap.String("job", rn.j.Id))
	return
}

func (rn *runner) Sync(ctx context.Context, job *models.Job) (err error) {
	logger := rn.logger

	_, err = rn.grpcClient.CreateJob(ctx, &models.CreateJobRequest{Job: job})
	if err != nil {
		return
	}

	_, err = rn.grpcClient.WaitJob(ctx, &models.WaitJobRequest{JobId: job.Id})
	if err != nil {
		logger.Error("wait job", zap.Error(err))
	}

	logger.Info("runner synced job",
		zap.String("job", job.Id))
	return
}

func (rn *runner) Finish(ctx context.Context, err error) error {
	logger := rn.logger

	logger.Info("runner finish job", zap.String("job", rn.j.Id))

	jp := &models.FinishJobRequest{
		JobId: rn.j.Id,
	}

	if err == nil {
		jp.Status = models.JobStatus_Succeed
	} else {
		jp.Status = models.JobStatus_Failed
		jp.Message = err.Error()
	}
	_, err = rn.grpcClient.FinishJob(ctx, jp)
	if err != nil {
		return err
	}
	return nil
}
