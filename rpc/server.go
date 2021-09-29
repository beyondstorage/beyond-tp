package rpc

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net"

	"github.com/beyondstorage/beyond-tp/proto"
	"github.com/beyondstorage/beyond-tp/task"
)

type Server struct {
	ts *task.Server

	addr string
	gs   *grpc.Server

	proto.UnimplementedAgentServer
}

type Config struct {
	Addr   string
	Task   *task.Server
	Logger *zap.Logger
}

func New(cfg *Config) *Server {
	s := &Server{
		ts:   cfg.Task,
		addr: cfg.Addr,
		gs: grpc.NewServer(grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(cfg.Logger),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		),
	}
	proto.RegisterAgentServer(s.gs, s)
	return s
}

func (s *Server) Serve(ctx context.Context) (err error) {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}

	err = s.gs.Serve(l)
	if err != nil {
		return fmt.Errorf("grpc serve: %w", err)
	}
	return nil
}

func (s *Server) Notify(stream proto.Agent_NotifyServer) error {
	for {
		n, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch n.Message {
		case task.ClientHello:
			has, err := s.ts.HasTask(stream.Context())
			if err != nil {
				return fmt.Errorf("has task: %w", err)
			}
			if has {
				err = stream.Send(&proto.Notification{Message: task.ServerTaskAvailable})
				if err != nil {
					return err
				}
			}
		default:
			panic(fmt.Errorf("invalid notification send from client: %s", n.Message))
		}
	}
}

func (s *Server) NextTask(ctx context.Context, req *proto.NextTaskRequest) (*proto.TaskReply, error) {
	t, err := s.ts.NextTask(ctx)
	if err != nil {
		return nil, fmt.Errorf("next task: %w", err)
	}
	if t == nil {
		return nil, status.Error(codes.NotFound, "no more tasks")
	}
	return task.FromTask(t), nil
}

func (s *Server) NextJob(ctx context.Context, req *proto.NextJobRequest) (*proto.JobReply, error) {
	j, err := s.ts.NextJob(ctx, req.TaskId)
	if err != nil {
		return nil, fmt.Errorf("next task: %w", err)
	}
	if j == nil {
		return nil, status.Error(codes.NotFound, "no more jobs")
	}
	return task.FromJob(j), nil
}

func (s *Server) CreateJob(ctx context.Context, req *proto.CreateJobRequest) (*proto.EmptyReply, error) {
	j := task.NewJob(req.TaskId, int(req.Type), req.Content)

	err := s.ts.InsertJob(ctx, j)
	if err != nil {
		return nil, fmt.Errorf("insert job: %w", err)
	}
	return &proto.EmptyReply{}, nil
}

func (s *Server) WaitJob(ctx context.Context, req *proto.WaitJobRequest) (*proto.EmptyReply, error) {
	err := s.ts.WaitJob(ctx, req.TaskId, req.JobId)
	if err != nil {
		return nil, fmt.Errorf("insert job: %w", err)
	}
	return &proto.EmptyReply{}, nil
}

func (s *Server) FinishJob(ctx context.Context, req *proto.FinishJobRequest) (*proto.EmptyReply, error) {
	// TODO: we need to should finish job with error.
	err := s.ts.DeleteJob(ctx, req.TaskId, req.JobId)
	if err != nil {
		return nil, fmt.Errorf("insert job: %w", err)
	}
	return &proto.EmptyReply{}, nil
}

func (s *Server) SetMeta(ctx context.Context, req *proto.MetaEntry) (*proto.EmptyReply, error) {
	err := s.ts.SetMeta(ctx, req.TaskId, req.JobId, req.Key, req.Value)
	if err != nil {
		return nil, fmt.Errorf("set meta: %w", err)
	}
	return &proto.EmptyReply{}, nil
}

func (s *Server) GetMeta(ctx context.Context, req *proto.MetaKey) (*proto.MetaValue, error) {
	v, err := s.ts.GetMeta(ctx, req.TaskId, req.JobId, req.Key)
	if err != nil {
		return nil, fmt.Errorf("set meta: %w", err)
	}
	return &proto.MetaValue{Value: v}, nil
}

func (s *Server) DeleteMeta(ctx context.Context, req *proto.MetaKey) (*proto.EmptyReply, error) {
	err := s.ts.DeleteMeta(ctx, req.TaskId, req.JobId, req.Key)
	if err != nil {
		return nil, fmt.Errorf("set meta: %w", err)
	}
	return &proto.EmptyReply{}, nil
}
