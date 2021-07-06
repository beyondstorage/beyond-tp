package task

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/beyondstorage/beyond-tp/models"
)

type Manager struct {
	cfg ManagerConfig

	logger     *zap.Logger
	db         *models.DB
	grpcServer *grpc.Server

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
	p.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	models.RegisterStaffServer(p.grpcServer, p)

	l, err := net.Listen("tcp", cfg.GrpcAddr())
	if err != nil {
		logger.Error("grpc server listen", zap.Error(err))
		return
	}

	go func() {
		err = p.grpcServer.Serve(l)
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

func (p *Manager) Poll(req *models.PollRequest, srv models.Staff_PollServer) (err error) {
	logger := p.logger

	for {
		reply := &models.PollReply{}

		taskId, err := p.db.NextStaffTask(nil, req.StaffId)
		if err != nil {
			logger.Error("next staff task", zap.Error(err))
			return err
		}

		// task_id == "" means there is no task for out staff.
		if taskId == "" {
			reply.Status = models.PollStatus_Empty

			err = srv.Send(reply)
			if err != nil {
				return err
			}

			// FIXME: we need to find a way to watch staff task changes.
			time.Sleep(60 * time.Second)
			return err
		}

		task, err := p.db.GetTask(taskId)
		if err != nil {
			logger.Error("get task", zap.Error(err))
			return err
		}

		reply.Status = models.PollStatus_Valid
		reply.Task = task

		err = srv.Send(reply)
		if err != nil {
			return err
		}

		logger.Info("polled task", zap.String("id", task.Id))

		err = p.db.DeleteStaffTask(nil, req.StaffId, taskId)
		if err != nil {
			return err
		}
	}
}

func (p *Manager) Finish(ctx context.Context, req *models.FinishRequest) (reply *models.FinishReply, err error) {
	logger := p.logger

	reply = &models.FinishReply{}

	t, err := p.db.GetTask(req.TaskId)
	if err != nil {
		logger.Error("get task", zap.String("id", req.TaskId))
		return
	}

	t.Status = models.TaskStatus_Finished

	err = p.db.UpdateTask(t)
	if err != nil {
		logger.Error("update task", zap.String("id", req.TaskId))
		return
	}
	return
}

func (p *Manager) Stop(ctx context.Context) (err error) {
	p.grpcServer.Stop()
	p.grpcServer = nil

	err = p.db.Close()
	if err != nil {
		p.logger.Error("close database", zap.Error(err))
	}
	p.db = nil

	return err
}
