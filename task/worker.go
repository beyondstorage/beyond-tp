package task

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aos-dev/go-storage/v3/types"
	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	natsproto "github.com/nats-io/nats.go/encoders/protobuf"
	"go.uber.org/zap"

	"github.com/aos-dev/dm/proto"
)

type Worker struct {
	id      string
	subject string

	queue    *nats.EncodedConn
	storages []types.Storager

	ctx    context.Context
	cond   *sync.Cond
	logger *zap.Logger
}

func NewWorker(ctx context.Context, addr, subject string, storages []types.Storager) (*Worker, error) {
	logger := zapcontext.From(ctx)

	w := &Worker{
		id:       uuid.NewString(),
		subject:  subject,
		storages: storages,

		ctx:    ctx,
		cond:   sync.NewCond(&sync.Mutex{}),
		logger: logger,
	}
	w.cond.L.Lock()

	// Connect to queue
	queueConn, err := nats.Connect(addr)
	if err != nil {
		return nil, err
	}
	w.queue, err = nats.NewEncodedConn(queueConn, natsproto.PROTOBUF_ENCODER)
	if err != nil {
		return nil, fmt.Errorf("nats encoded connect: %w", err)
	}

	logger.Info("worker has been setup", zap.String("id", w.id))
	return w, nil
}

func (w *Worker) clockin() {
	w.logger.Info("worker start clockin", zap.String("id", w.id))

	reply := &proto.ClockinReply{}

	for {
		err := w.queue.RequestWithContext(w.ctx, SubjectClockin(w.subject),
			&proto.ClockinRequest{}, reply)
		if err != nil && errors.Is(err, nats.ErrNoResponders) {
			time.Sleep(25 * time.Millisecond)
			continue
		}
		if err != nil {
			w.logger.Error("worker clockin", zap.String("id", w.id), zap.Error(err))
			return
		}
		break
	}
}

func (w *Worker) clockout() {
	w.logger.Info("worker start waiting for clockout", zap.String("id", w.id))

	_, err := w.queue.Subscribe(SubjectClockoutNotify(w.subject),
		func(subject, reply string, req *proto.ClockoutRequest) {
			err := w.queue.Publish(reply, &proto.Acknowledgement{})
			if err != nil {
				w.logger.Error("publish ack", zap.Error(err))
				return
			}

			w.cond.Signal()
		})
	if err != nil {
		return
	}
}

func (w *Worker) Handle(ctx context.Context) (err error) {
	go w.clockout()

	// Worker must setup before clockin.
	_, err = w.queue.QueueSubscribe(w.subject, w.subject,
		func(subject, reply string, job *proto.Job) {
			go func() {
				rn, err := NewRunner(w, job)
				if err != nil {
					w.logger.Error("create new runner", zap.Error(err))
					return
				}
				rn.Handle(reply)
			}()
		})
	if err != nil {
		return fmt.Errorf("nats subscribe: %w", err)
	}

	// TODO: we can clockout directly if the task has been finished.
	w.clockin()

	w.cond.Wait()
	return
}

func HandleAsWorker(ctx context.Context, addr, subject string, storages []types.Storager) (err error) {
	w, err := NewWorker(ctx, addr, subject, storages)
	if err != nil {
		return
	}

	return w.Handle(ctx)
}
