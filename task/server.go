package task

import (
	"context"
	"net"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/beyondstorage/beyond-tp/models"
)

type Server struct {
	cfg ServerConfig

	logger     *zap.Logger
	db         *models.DB
	grpcServer *grpc.Server

	models.UnimplementedAgentServer
}

type ServerConfig struct {
	GrpcAddr string

	DatabasePath string
}

func NewServer(ctx context.Context, cfg ServerConfig) (s *Server, err error) {
	logger := zapcontext.From(ctx)

	s.db, err = models.NewDB(s.cfg.DatabasePath, logger)
	if err != nil {
		logger.Error("create db",
			zap.String("path", s.cfg.DatabasePath), zap.Error(err))
		return
	}

	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(
		middleware.ChainUnaryServer(
			grpczap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	models.RegisterAgentServer(s.grpcServer, s)

	l, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		logger.Error("grpc server listen", zap.Error(err))
		return
	}

	go func() {
		err = s.grpcServer.Serve(l)
		if err != nil {
			logger.Error("grpc server serve", zap.Error(err))
			return
		}
	}()

	logger.Info("server has been setup", zap.String("addr", cfg.GrpcAddr))
	return s, nil
}

func (s *Server) PollJob(ctx context.Context, req *models.PollJobRequest, opts ...grpc.CallOption) (models.Agent_PollJobClient, error) {
	logger := s.logger

	for j := range s.jobCh {
		err = srv.Send(&models.PollJobReply{
			Status: models.PollJobStatus_Valid,
			Job:    j,
		})
		if err != nil {
			logger.Error("send next job", zap.Error(err))
			return
		}
	}

	// If job channel has been closed, that means no more new jobs will be added.
	err = srv.Send(&models.PollJobReply{
		Status: models.PollJobStatus_Terminated,
	})
	if err != nil {
		logger.Error("send terminate job", zap.Error(err))
	}
	return
}

func (s *Server) CreateJob(ctx context.Context, req *models.CreateJobRequest, opts ...grpc.CallOption) (reply *models.CreateJobReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.CreateJobReply{
		Status: 0,
	}

	err = s.db.InsertJob(req.Job)
	if err != nil {
		logger.Error("insert job", zap.Error(err))
		return
	}
	return
}

func (s *Server) WaitJob(ctx context.Context, req *models.WaitJobRequest, opts ...grpc.CallOption) (reply *models.WaitJobReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.WaitJobReply{
		Status: 0,
	}

	err = s.db.WaitJob(ctx, req.JobId)
	if err != nil {
		logger.Error("wait job", zap.Error(err))
		return
	}
	return
}

func (s *Server) FinishJob(ctx context.Context, req *models.FinishJobRequest, opts ...grpc.CallOption) (reply *models.FinishJobReply, err error) {
	logger := zapcontext.From(ctx)

	reply := &models.FinishJobReply{}

	err = s.db.DeleteJob(ctx, req.JobId)
	if err != nil {
		logger.Error("delete job", zap.Error(err))
		return
	}
	logger.Debug("job finished", zap.String("id", req.JobId), zap.String("status", req.Status.String()),
		zap.String("root job", s.rootJobId))

	if req.JobId == s.rootJobId {
		logger.Debug("root job finished", zap.String("id", req.JobId))

		close(s.doneCh)
		close(s.jobCh)
	}
	return
}

func (s *Server) GetJobMetadata(ctx context.Context, req *models.GetJobMetadataRequest) (
	*models.GetJobMetadataReply, error) {
	res, err := s.db.GetJobMetadata(req.JobId)
	if err != nil {
		return nil, err
	}
	return &models.GetJobMetadataReply{Metadata: res}, nil
}

func (s *Server) SetJobMetadata(ctx context.Context, req *models.SetJobMetadataRequest) (
	*models.SetJobMetadataReply, error) {
	err := s.db.SetJobMetadata(req.JobId, req.Metadata)
	if err != nil {
		return nil, err
	}
	return &models.SetJobMetadataReply{}, nil
}

func (s *Server) DeleteJobMetadata(ctx context.Context, req *models.DeleteJobMetadataRequest) (
	*models.DeleteJobMetadataReply, error) {
	err := s.db.DeleteJobMetadata(req.JobId)
	if err != nil {
		return nil, err
	}
	return &models.DeleteJobMetadataReply{}, nil
}
