package task

import (
	"context"
	"fmt"
	"github.com/aos-dev/go-storage/v3/types"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"time"

	"github.com/aos-dev/go-toolbox/natszap"
	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/google/uuid"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	natsproto "github.com/nats-io/nats.go/encoders/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/aos-dev/dm/proto"
)

type Staff struct {
	id   string
	addr string
	cfg  StaffConfig

	logger *zap.Logger
	ctx    context.Context

	grpcClient proto.StaffClient
	queueSrv   *server.Server

	queue *nats.EncodedConn
	sub   *nats.Subscription
}

type StaffConfig struct {
	Host string

	ManagerAddr string
}

func NewStaff(ctx context.Context, cfg StaffConfig) (s *Staff, err error) {
	logger := zapcontext.From(ctx)

	s = &Staff{
		id:  uuid.New().String(),
		cfg: cfg,

		ctx:    ctx,
		logger: logger,
	}

	// Setup NATS server.
	s.queueSrv, err = server.NewServer(&server.Options{
		Host: cfg.Host,
		Port: server.RANDOM_PORT,
	})
	if err != nil {
		return
	}

	go func() {
		s.queueSrv.SetLoggerV2(natszap.NewLog(logger), false, false, false)

		s.queueSrv.Start()
	}()

	if !s.queueSrv.ReadyForConnections(10 * time.Second) {
		panic(fmt.Errorf("server start too slow"))
	}
	s.addr = s.queueSrv.ClientURL()
	return
}

// Connect will connect to portal task queue.
func (s *Staff) Connect(ctx context.Context) (err error) {
	logger := s.logger

	// FIXME: we need to use ssl/tls to encrypt our channel.
	grpcConn, err := grpc.DialContext(ctx, s.cfg.ManagerAddr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_zap.UnaryClientInterceptor(logger)),
	)
	if err != nil {
		return
	}
	s.grpcClient = proto.NewStaffClient(grpcConn)

	reply, err := s.grpcClient.Register(ctx, &proto.RegisterRequest{
		Id:   s.id,
		Addr: s.addr,
	})
	if err != nil {
		return
	}

	logger.Info("connect to task queue",
		zap.String("addr", reply.Addr),
		zap.String("subject", reply.Subject))

	conn, err := nats.Connect(reply.Addr)
	if err != nil {
		return
	}

	s.queue, err = nats.NewEncodedConn(conn, natsproto.PROTOBUF_ENCODER)
	if err != nil {
		return
	}
	s.sub, err = s.queue.Subscribe(reply.Subject,
		func(subject, reply string, task *proto.Task) {
			s.logger.Info("start handle task",
				zap.String("subject", subject),
				zap.String("id", task.Id),
				zap.String("staff_id", s.id))

			go s.Handle(reply, task)
		})
	if err != nil {
		return
	}
	return nil
}

// Handle will create a new agent to handle task.
func (s *Staff) Handle(reply string, task *proto.Task) {
	// Parse storage
	storages := make([]types.Storager, 0)
	for _, ep := range task.Endpoints {
		store, err := ep.ParseStorager()
		if err != nil {
			s.logger.Error("parse storager", zap.Error(err))
			return
		}
		storages = append(storages, store)
	}

	// Send upgrade
	electReply, err := s.grpcClient.Elect(s.ctx, &proto.ElectRequest{
		StaffId: s.id,
		TaskId:  task.Id,
	})
	if err != nil {
		s.logger.Error("staff elect", zap.String("id", s.id), zap.Error(err))
		return
	}

	tr := &proto.TaskReply{Id: task.Id, StaffId: s.id}

	if electReply.LeaderId == s.id {
		err = HandleAsLeader(s.ctx, electReply.Addr, electReply.Subject, electReply.WorkerIds, task.Job)
	} else {
		err = HandleAsWorker(s.ctx, electReply.Addr, electReply.Subject, storages)
	}
	if err == nil {
		tr.Status = JobStatusSucceed
	} else {
		tr.Status = JobStatusFailed
		tr.Message = fmt.Sprintf("task handle: %v", err)
	}
	err = nil

	err = s.queue.Publish(reply, tr)
	if err != nil {
		s.logger.Error("staff reply",
			zap.String("id", s.id), zap.Error(err))
	}
}
