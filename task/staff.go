package task

import (
	"context"
	"fmt"
	"github.com/aos-dev/go-storage/v3/types"
	"github.com/aos-dev/go-toolbox/zapcontext"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"net"
	"path/filepath"
	"sync"

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
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.DefaultConfig,
		}),
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

	// TODO: we will support multi tasks later, for now, we only support one task.
	var currentTaskId string

	srv, err := s.grpcClient.Poll(ctx, &models.PollRequest{StaffId: s.id})
	if err != nil {
		logger.Error("receive next task", zap.Error(err))
		return err
	}

	for {
		t, err := srv.Recv()
		if err != nil {
			logger.Error("receive next task", zap.Error(err))
			return err
		}
		// If status is terminated, we should exit the staff.
		if t.Status == models.PollStatus_Terminated {
			return nil
		}
		// If status is empty, we should wait next tick.
		if t.Status == models.PollStatus_Empty {
			continue
		}
		// If current task id equals to polled task, just start next round trip.
		if t.Task.Id == currentTaskId {
			continue
		} else {
			currentTaskId = t.Task.Id
		}

		l, err := net.Listen("tcp", s.cfg.Host+":0")
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
		logger.Debug("elect leader", zap.String("staff_id", reply.LeaderId))

		if reply.LeaderId == s.id {
			dp := filepath.Join(s.cfg.DataPath, t.Task.Id)
			job := mapTaskToRootJob(t.Task)

			cond := sync.NewCond(&sync.Mutex{})
			cond.L.Lock()

			go func() {
				cond.Wait()

				err = s.FinishTask(ctx, t.Task.Id)
				if err != nil {
					logger.Error("finish task", zap.String("id", t.Task.Id))
				}
				logger.Info("finish task", zap.String("id", t.Task.Id))
			}()
			go HandleAsLeader(ctx, l, dp, cond, job)
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

func (s *Staff) FinishTask(ctx context.Context, taskId string) (err error) {
	_, err = s.grpcClient.Finish(ctx, &models.FinishRequest{
		TaskId: taskId,
	})
	return
}

func mapTaskToRootJob(t *models.Task) *models.Job {
	j := &models.Job{
		Id: uuid.NewString(),
	}

	var innerJob protobuf.Message
	switch t.Type {
	case models.TaskType_CopyDir:
		j.Type = models.JobType_CopyDir
		innerJob = &models.CopyDirJob{
			Src:       0,
			Dst:       1,
			SrcPath:   "",
			DstPath:   "",
			Recursive: true,
		}
	default:
		panic(fmt.Errorf("task type %s is not supported", t.Type.String()))
	}

	bs, err := protobuf.Marshal(innerJob)
	if err != nil {
		panic(fmt.Errorf("protobuf marshal: %v", err))
	}
	j.Content = bs
	return j
}
