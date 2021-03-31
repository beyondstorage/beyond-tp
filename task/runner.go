package task

import (
	"context"
	"fmt"
	"sync"

	"github.com/aos-dev/go-storage/v3/types"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/proto"
)

type Runner struct {
	agent *Worker
	j     *proto.Job

	queue    *nats.EncodedConn
	subject  string // All runner will share the same task subject
	storages []types.Storager

	wg     *sync.WaitGroup
	sub    *nats.Subscription
	logger *zap.Logger
}

func NewRunner(a *Worker, j *proto.Job) (*Runner, error) {
	rn := &Runner{
		j:     j,
		agent: a,

		queue:    a.queue,
		subject:  a.subject, // Copy task subject from agent.
		storages: a.storages,
		wg:       &sync.WaitGroup{},
		logger:   a.logger,
	}

	var err error
	// Wait for all JobReply sending to the reply subject.
	rn.sub, err = rn.queue.Subscribe(SubjectJobReply(rn.j.Id), func(job *proto.JobReply) {
		defer rn.wg.Done()

		switch job.Status {
		case JobStatusSucceed:
			rn.logger.Info("job succeed",
				zap.String("id", job.Id),
				zap.String("job", rn.j.Id))
		default:
			rn.logger.Error("job failed",
				zap.String("id", job.Id),
				zap.String("job", rn.j.Id),
				zap.String("error", job.Message),
			)
		}
	})
	if err != nil {
		rn.logger.Error("runner subscribe job reply", zap.Error(err))
		return nil, err
	}
	return rn, nil
}

func (rn *Runner) Handle(reply string) {
	rn.logger.Info("runner start job",
		zap.String("id", rn.j.Id),
		zap.Uint32("type", rn.j.Type))

	ctx := context.Background()

	var fn func(ctx context.Context, msg protobuf.Message) error
	var t protobuf.Message

	switch rn.j.Type {
	case TypeCopyDir:
		t = &proto.CopyDir{}
		fn = rn.HandleCopyDir
	case TypeCopyFile:
		t = &proto.CopyFile{}
		fn = rn.HandleCopyFile
	case TypeCopySingleFile:
		t = &proto.CopySingleFile{}
		fn = rn.HandleCopySingleFile
	case TypeCopyMultipartFile:
		t = &proto.CopyMultipartFile{}
		fn = rn.HandleCopyMultipartFile
	case TypeCopyMultipart:
		t = &proto.CopyMultipart{}
		fn = rn.HandleCopyMultipart
	default:
		panic("not support job type")
	}

	err := protobuf.Unmarshal(rn.j.Content, t)
	if err != nil {
		panic(fmt.Errorf("job unmarshal, %w", err))
	}

	err = fn(ctx, t)

	// Send JobReply after the job has been handled.
	err = rn.Finish(ctx, reply, err)
	if err != nil {
		rn.logger.Error("runner finish", zap.Error(err))
	}
}

func (rn *Runner) Async(ctx context.Context, job *proto.Job) (err error) {
	logger := rn.logger

	rn.wg.Add(1)

	// Publish new job with the specific reply subject on the task subject.
	// After this job finished, the runner will send a JobReply to the reply subject.
	err = rn.queue.PublishRequest(rn.subject, SubjectJobReply(rn.j.Id), job)
	if err != nil {
		logger.Error("runner publish", zap.Error(err))
		return fmt.Errorf("nats publish: %w", err)
	}

	logger.Info("runner publish async job",
		zap.String("subject", rn.subject),
		zap.String("job", rn.j.Id),
		zap.String("id", job.Id))
	return
}

func (rn *Runner) Await(ctx context.Context) (err error) {
	rn.logger.Info("runner start await job",
		zap.String("job", rn.j.Id))

	rn.wg.Wait()

	rn.logger.Info("runner finish await job",
		zap.String("job", rn.j.Id))
	return
}

func (rn *Runner) Sync(ctx context.Context, job *proto.Job) (err error) {
	logger := rn.logger

	var reply proto.JobReply

	logger.Info("runner start sync job",
		zap.String("subject", rn.subject),
		zap.String("job", rn.j.Id),
		zap.String("id", job.Id))

	// NATS provides the builtin request-response style API, so that we don't need to
	// care about the reply id.
	err = rn.queue.RequestWithContext(ctx, rn.subject, job, &reply)
	if err != nil {
		logger.Error("runner request", zap.Error(err))
		return fmt.Errorf("nats request: %w", err)
	}

	if reply.Status != JobStatusSucceed {
		logger.Error("job synced",
			zap.String("job", reply.Id),
			zap.String("error", reply.Message))
		return fmt.Errorf("job failed: %v", reply.Message)
	}

	logger.Info("runner synced job",
		zap.String("subject", rn.subject),
		zap.String("job", rn.j.Id),
		zap.String("id", job.Id))
	return
}

func (rn *Runner) Finish(ctx context.Context, reply string, err error) error {
	logger := rn.logger

	logger.Info("runner reply",
		zap.String("job", rn.j.Id),
		zap.String("reply", reply))

	jp := &proto.JobReply{
		Id: rn.j.Id, // Make sure JobReply sends to the parent job.
	}

	if err == nil {
		jp.Status = JobStatusSucceed
	} else {
		jp.Status = JobStatusFailed
		jp.Message = err.Error()
	}
	return rn.queue.Publish(reply, jp)
}
