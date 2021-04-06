package task

import (
	"context"
	"sync"

	"github.com/aos-dev/go-storage/v3/types"
	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/google/uuid"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"

	"github.com/aos-dev/dm/models"
)

type Worker struct {
	id string

	grpcClient models.WorkerClient
	storages   []types.Storager

	ctx    context.Context
	logger *zap.Logger
}

func NewWorker(ctx context.Context, addr string, storages []types.Storager) (w *Worker, err error) {
	logger := zapcontext.From(ctx)

	w = &Worker{
		id:       uuid.NewString(),
		storages: storages,

		ctx:    ctx,
		logger: logger,
	}

	grpcConn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.DefaultConfig,
		}),
		grpc.WithUnaryInterceptor(grpc_zap.UnaryClientInterceptor(logger)),
	)
	if err != nil {
		logger.Error("dial manager", zap.Error(err))
		return nil, err
	}
	w.grpcClient = models.NewWorkerClient(grpcConn)

	logger.Info("worker has been setup", zap.String("id", w.id))
	return w, nil
}

func (w *Worker) Serve(ctx context.Context) (err error) {
	logger := zapcontext.From(ctx)

	jc, err := w.grpcClient.PollJob(ctx, &models.PollJobRequest{})
	if err != nil {
		logger.Error("next job", zap.Error(err))
		return err
	}

	for {
		j, err := jc.Recv()
		if err != nil {
			logger.Error("receive next job", zap.Error(err))
			return err
		}
		if j.Status == models.PollJobStatus_InvalidPollJobStatus {
			panic("invalid poll job status")
		}
		if j.Status == models.PollJobStatus_Terminated {
			return nil
		}

		go func() {
			rn, err := newRunner(w, j.Job)
			if err != nil {
				w.logger.Error("create new runner", zap.Error(err))
				return
			}
			rn.Handle()
		}()
	}
}

func HandleAsWorker(ctx context.Context, addr string, cond *sync.Cond, storages []types.Storager) {
	logger := zapcontext.From(ctx)

	w, err := NewWorker(ctx, addr, storages)
	if err != nil {
		return
	}

	err = w.Serve(ctx)
	if err != nil {
		logger.Error("worker serve", zap.Error(err))
		return
	}

	cond.Signal()
	return
}
