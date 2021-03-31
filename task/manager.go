package task

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/aos-dev/go-toolbox/natszap"
	"github.com/aos-dev/go-toolbox/zapcontext"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	natsproto "github.com/nats-io/nats.go/encoders/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/aos-dev/dm/proto"
)

type Manager struct {
	queue *nats.EncodedConn

	staffLock  sync.RWMutex
	staffIds   []string
	staffAddrs map[string]string

	taskLock sync.RWMutex
	tasks    map[string]*taskMeta

	config ManagerConfig

	proto.UnimplementedStaffServer
}

type taskMeta struct {
	sub *nats.Subscription
	wg  *sync.WaitGroup
}

type ManagerConfig struct {
	Host     string
	GrpcPort int

	// Queue related config.
	QueuePort int
}

func (p ManagerConfig) GrpcAddr() string {
	return fmt.Sprintf("%s:%d", p.Host, p.GrpcPort)
}

func (p ManagerConfig) QueueAddr() string {
	return fmt.Sprintf("%s:%d", p.Host, p.QueuePort)
}

func NewManager(ctx context.Context, cfg ManagerConfig) (p *Manager, err error) {
	logger := zapcontext.From(ctx)

	p = &Manager{
		config: cfg,

		staffAddrs: make(map[string]string),
		tasks:      make(map[string]*taskMeta),
	}

	// Setup grpc server.
	grpcSrv := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	proto.RegisterStaffServer(grpcSrv, p)
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

	// Setup queue server.
	srv, err := server.NewServer(&server.Options{
		Host: cfg.Host,
		Port: cfg.QueuePort,
	})
	if err != nil {
		logger.Error("create nats server", zap.Error(err))
		return
	}

	go func() {
		srv.SetLoggerV2(natszap.NewLog(logger), false, false, false)

		srv.Start()
	}()

	if !srv.ReadyForConnections(10 * time.Second) {
		panic(fmt.Errorf("server start too slow"))
	}

	conn, err := nats.Connect(srv.ClientURL())
	if err != nil {
		logger.Error("connect nats queue",
			zap.String("addr", srv.ClientURL()), zap.Error(err))
		return
	}
	p.queue, err = nats.NewEncodedConn(conn, natsproto.PROTOBUF_ENCODER)
	if err != nil {
		logger.Error("connect encoded nats queue",
			zap.String("addr", srv.ClientURL()), zap.Error(err))
		return
	}

	return p, nil
}

func (p *Manager) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.RegisterReply, error) {
	_ = zapcontext.From(ctx)

	p.staffLock.Lock()
	defer p.staffLock.Unlock()
	p.staffIds = append(p.staffIds, request.Id)
	p.staffAddrs[request.Id] = request.Addr

	return &proto.RegisterReply{
		Addr:    p.config.QueueAddr(),
		Subject: SubjectTasks(),
	}, nil
}

func (p *Manager) Elect(ctx context.Context, request *proto.ElectRequest) (*proto.ElectReply, error) {
	_ = zapcontext.From(ctx)

	p.staffLock.RLock()
	defer p.staffLock.RUnlock()

	return &proto.ElectReply{
		Addr:      p.staffAddrs[p.staffIds[0]],
		Subject:   SubjectTask(request.TaskId),
		LeaderId:  p.staffIds[0],
		WorkerIds: p.staffIds[1:],
	}, nil
}

// Publish will publish a task on "tasks" queue.
func (p *Manager) Publish(ctx context.Context, task *proto.Task) (err error) {
	logger := zapcontext.From(ctx)

	// TODO: We need to maintain all tasks in db maybe.
	logger.Info("manager publish task", zap.String("id", task.Id))

	tm := &taskMeta{
		wg: &sync.WaitGroup{},
	}

	tm.wg.Add(len(p.staffIds))
	// Subscribe task reply before we publish our request.
	tm.sub, err = p.queue.Subscribe(SubjectTaskReply(task.Id), func(tr *proto.TaskReply) {
		defer tm.wg.Done()

		switch tr.Status {
		case JobStatusSucceed:
			logger.Info("task succeed",
				zap.String("id", tr.Id),
				zap.String("staff_id", tr.StaffId))
		default:
			logger.Error("task failed",
				zap.String("id", tr.Id),
				zap.String("staff_id", tr.StaffId),
				zap.String("error", tr.Message),
			)
		}
	})
	if err != nil {
		return err
	}

	p.taskLock.Lock()
	p.tasks[task.Id] = tm
	p.taskLock.Unlock()

	// Send task on tasks and wait for reply.
	err = p.queue.PublishRequest(SubjectTasks(), SubjectTaskReply(task.Id), task)
	if err != nil {
		return
	}
	return
}

// Wait will wait for all staffAddrs' replies on specific task.
func (p *Manager) Wait(ctx context.Context, task *proto.Task) (err error) {
	logger := zapcontext.From(ctx)

	p.taskLock.RLock()
	tm := p.tasks[task.Id]
	p.taskLock.RUnlock()

	logger.Info("manager wait task to be finished",
		zap.String("id", task.Id))

	tm.wg.Wait()
	err = tm.sub.Unsubscribe()
	if err != nil {
		return
	}

	logger.Info("manager finished task", zap.String("id", task.Id))

	p.taskLock.Lock()
	delete(p.tasks, task.Id)
	p.taskLock.Unlock()
	return
}
