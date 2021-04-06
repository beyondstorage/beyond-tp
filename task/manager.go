package task

import (
	"context"
	"fmt"
	"github.com/aos-dev/go-toolbox/zapcontext"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"

	"github.com/aos-dev/dm/models"
)

type Manager struct {
	cfg ManagerConfig

	logger *zap.Logger
	db     *models.DB

	models.UnimplementedStaffServer
}

func (p *Manager) DB() *models.DB {
	return p.db
}

type ManagerConfig struct {
	Host     string
	GrpcPort int

	DatabasePath string
}

func (p ManagerConfig) GrpcAddr() string {
	return fmt.Sprintf("%s:%d", p.Host, p.GrpcPort)
}

func NewManager(ctx context.Context, cfg ManagerConfig) (p *Manager, err error) {
	logger := zapcontext.From(ctx)

	p = &Manager{
		cfg: cfg,

		logger: logger,
	}

	p.db, err = models.NewDB(p.cfg.DatabasePath, logger)
	if err != nil {
		logger.Error("create db", zap.String("path", p.cfg.DatabasePath), zap.Error(err))
		return
	}

	// Setup grpc server.
	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	models.RegisterStaffServer(grpcSrv, p)

	l, err := net.Listen("tcp", cfg.GrpcAddr())
	if err != nil {
		logger.Error("grpc server listen", zap.Error(err))
		return
	}

	go func() {
		err = grpcSrv.Serve(l)
		if err != nil {
			logger.Error("grpc server serve", zap.Error(err))
			return
		}
	}()

	logger.Info("manager has been setup", zap.String("addr", cfg.GrpcAddr()))
	return p, nil
}

func (p *Manager) Register(ctx context.Context, req *models.RegisterRequest) (reply *models.RegisterReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.RegisterReply{}

	_, err = p.db.CreateStaff(req.StaffId)
	if err != nil {
		logger.Error("create staff", zap.Error(err))
		return
	}

	return reply, nil
}

func (p *Manager) Elect(ctx context.Context, req *models.ElectRequest) (reply *models.ElectReply, err error) {
	logger := zapcontext.From(ctx)

	reply = &models.ElectReply{}

	id, addr, err := p.db.ElectTaskLeader(req.TaskId, req.StaffId, req.StaffAddr)
	if err != nil {
		logger.Error("elect task leader", zap.Error(err))
		return
	}

	return &models.ElectReply{
		LeaderId:   id,
		LeaderAddr: addr,
	}, nil
}

func (p *Manager) Poll(ctx context.Context, req *models.PollRequest) (reply *models.PollReply, err error) {
	logger := p.logger

	reply = &models.PollReply{}

	// TODO: we need to add status for staff task.
	taskId, err := p.db.NextStaffTask(nil, req.StaffId)
	if err != nil {
		logger.Error("next staff task", zap.Error(err))
		return
	}
	logger.Debug("got task",
		zap.String("id", taskId),
		zap.String("staff_id", req.StaffId))

	if taskId == "" {
		reply.Status = models.PollStatus_Empty
		return
	}

	task, err := p.db.GetTask(taskId)
	if err != nil {
		logger.Error("get task", zap.Error(err))
		return
	}

	reply.Status = models.PollStatus_Valid
	reply.Task = task

	logger.Info("polled task", zap.String("id", task.Id))
	return
}

func (p *Manager) Finish(ctx context.Context, req *models.FinishRequest) (reply *models.FinishReply, err error) {
	logger := p.logger

	reply = &models.FinishReply{}

	err = p.db.DeleteStaffTask(nil, req.StaffId, req.TaskId)
	if err != nil {
		logger.Error("delete staff task", zap.Error(err))
		return
	}
	return
}
