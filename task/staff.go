package task

import (
	"context"
	"net"
	"path/filepath"

	"github.com/aos-dev/go-storage/v3/types"
	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/google/uuid"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/aos-dev/dm/models"
)

type Staff struct {
	id  string
	cfg StaffConfig

	ctx        context.Context
	logger     *zap.Logger
	grpcClient models.StaffClient
}

type StaffConfig struct {
	Host     string
	DataPath string

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
	return
}

// Connect will connect to portal task queue.
func (s *Staff) Start(ctx context.Context) (err error) {
	logger := s.logger

	// FIXME: we need to use ssl/tls to encrypt our channel.
	grpcConn, err := grpc.DialContext(ctx, s.cfg.ManagerAddr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_zap.UnaryClientInterceptor(logger)),
	)
	if err != nil {
		logger.Error("dial manager", zap.Error(err))
		return
	}
	s.grpcClient = models.NewStaffClient(grpcConn)

	_, err = s.grpcClient.Register(ctx, &models.RegisterRequest{
		StaffId: s.id,
	})
	if err != nil {
		logger.Error("register", zap.Error(err))
		return
	}

	tc, err := s.grpcClient.NextTask(ctx, &models.NextTaskRequest{StaffId: s.id})
	if err != nil {
		logger.Error("next task", zap.Error(err))
		return err
	}

	for {
		t, err := tc.Recv()
		if err != nil {
			logger.Error("receive next task", zap.Error(err))
			return err
		}

		l, err := net.Listen("tcp", s.cfg.Host)
		if err != nil {
			logger.Error("grpc server listen", zap.Error(err))
			return err
		}

		reply, err := s.grpcClient.Elect(ctx, &models.ElectRequest{
			StaffId:   s.id,
			StaffAddr: l.Addr().String(),
			TaskId:    t.Task.Id,
		})
		if err != nil {
			logger.Error("grpc elect", zap.Error(err))
			return err
		}

		if reply.LeaderId == s.id {
			dp := filepath.Join(s.cfg.DataPath, t.Task.Id)
			go HandleAsLeader(ctx, l, dp)
		} else {
			// Close listener as we don't need it anymore.
			l.Close()

			sts := make([]types.Storager, 0, len(t.Task.Storages))
			for _, v := range t.Task.Storages {
				store, err := models.FormatStorage(v)
				if err != nil {
					logger.Error("format storage", zap.Error(err))
					return err
				}

				sts = append(sts, store)
			}

			go HandleAsWorker(ctx, reply.LeaderAddr, sts)
		}
	}
}
