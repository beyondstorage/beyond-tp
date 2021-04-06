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
	taskCh chan *models.Task

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
		cfg:    cfg,
		taskCh: make(chan *models.Task),

		logger: logger,
	}

	p.db, err = models.NewDB(p.cfg.DatabasePath)
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
	go func() {
		l, err := net.Listen("tcp", cfg.GrpcAddr())
		if err != nil {
			logger.Error("grpc server listen", zap.Error(err))
			return
		}
		err = grpcSrv.Serve(l)
		if err != nil {
			logger.Error("grpc server serve", zap.Error(err))
			return
		}
	}()

	return p, nil
}

func (p *Manager) Serve(ctx context.Context) {
	logger := zapcontext.From(ctx)

	defer close(p.taskCh)

	err := p.db.SubscribeTask(ctx, func(t *models.Task) {
		if t.Status == models.TaskStatus_Ready {
			p.taskCh <- t
		}
	})
	if err != nil {
		logger.Error("subscribe task", zap.Error(err))
		return
	}
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

func (p *Manager) NextTask(req *models.NextTaskRequest, srv models.Staff_NextTaskServer) (err error) {
	logger := p.logger

	for t := range p.taskCh {
		err = srv.Send(&models.NextTaskReply{
			Status: 0,
			Task:   t,
		})
		if err != nil {
			logger.Error("send next task", zap.Error(err))
			return
		}
	}
	return
}
