package task

import (
	"context"
	"net"

	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/aos-dev/dm/models"
)

type Leader struct {
	id string

	db      *models.DB
	jobCh   chan *models.Job
	grpcSrv models.WorkerServer

	logger *zap.Logger

	models.UnimplementedWorkerServer
}

func NewLeader(ctx context.Context,
	nl net.Listener,
	databasePath string,
) (l *Leader, err error) {
	logger := zapcontext.From(ctx)

	l = &Leader{
		id: uuid.NewString(),

		logger: logger,
	}

	l.db, err = models.NewDB(databasePath)
	if err != nil {
		logger.Error("create db", zap.Error(err))
		return
	}

	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	models.RegisterWorkerServer(grpcSrv, l)
	go func() {
		err = grpcSrv.Serve(nl)
		if err != nil {
			logger.Error("grpc server serve", zap.Error(err))
			return
		}
	}()

	logger.Info("leader has been setup", zap.String("id", l.id))
	return
}

func (l *Leader) Serve(ctx context.Context) (err error) {
	logger := zapcontext.From(ctx)

	defer close(l.jobCh)

	err = l.db.SubscribeJob(ctx, func(t *models.Job) {
		l.jobCh <- t
	})
	if err != nil {
		logger.Error("subscribe job", zap.Error(err))
		return
	}
	return
}

func (l *Leader) NextJob(request *models.NextJobRequest, srv models.Worker_NextJobServer) (err error) {
	logger := l.logger

	for j := range l.jobCh {
		err = srv.Send(&models.NextJobReply{
			Status: 0,
			Job:    j,
		})
		if err != nil {
			logger.Error("send next job", zap.Error(err))
			return
		}
	}
	return
}

func (l *Leader) CreateJob(ctx context.Context, req *models.CreateJobRequest) (reply *models.CreateJobReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.CreateJobReply{
		Status: 0,
	}

	err = l.db.InsertJob(req.Job)
	if err != nil {
		logger.Error("insert job", zap.Error(err))
		return
	}
	return
}

func (l *Leader) WaitJob(ctx context.Context, req *models.WaitJobRequest) (reply *models.WaitJobReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.WaitJobReply{
		Status: 0,
	}

	err = l.db.WaitJob(ctx, req.JobId)
	if err != nil {
		logger.Error("wait job", zap.Error(err))
		return
	}
	return
}

func (l *Leader) FinishJob(ctx context.Context, req *models.FinishJobRequest) (reply *models.FinishJobReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.FinishJobReply{}

	err = l.db.DeleteJob(ctx, req.JobId)
	if err != nil {
		logger.Error("delete job", zap.Error(err))
		return
	}
	return
}

func HandleAsLeader(ctx context.Context, nl net.Listener, dp string) {
	logger := zapcontext.From(ctx)

	l, err := NewLeader(ctx, nl, dp)
	if err != nil {
		logger.Error("create new leader", zap.Error(err))
		return
	}

	err = l.Serve(ctx)
	if err != nil {
		logger.Error("leader serve", zap.Error(err))
	}
}
