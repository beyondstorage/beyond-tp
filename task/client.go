package task

import (
	"context"
	"fmt"
	"io"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/beyondstorage/beyond-tp/proto"
)

type Client struct {
	pool   *ants.Pool
	gc     proto.AgentClient // gRPC Client
	logger *zap.Logger
	errch  chan error
}

type ClientConfig struct {
	Addr   string
	Logger *zap.Logger
}

func NewClient(cc *ClientConfig) (c *Client, err error) {
	c = &Client{
		logger: cc.Logger,
		errch:  make(chan error),
	}

	// Set default worker pool
	c.pool, _ = ants.NewPool(2)

	// Setup grpc client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, cc.Addr,
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.DefaultConfig,
		}),
		grpc.WithUnaryInterceptor(grpc_zap.UnaryClientInterceptor(c.logger)))
	if err != nil {
		return nil, fmt.Errorf("dial %s: %v", cc.Addr, err)
	}
	c.gc = proto.NewAgentClient(conn)
	return
}

func (c *Client) Error() chan error {
	return c.errch
}

func (c *Client) Start(ctx context.Context) (err error) {
	nc, err := c.gc.Notify(ctx)
	if err != nil {
		return fmt.Errorf("notify: %w", err)
	}

	for {
		n, err := nc.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch n.Message {
		case ServerTaskAvailable:
			err = c.NextTask(ctx)
			if err != nil {
				return err
			}
		default:
			panic(fmt.Errorf("invalid notification send from server: %s", n.Message))
		}
	}
}

func (c *Client) NextTask(ctx context.Context) (err error) {
	tr, err := c.gc.NextTask(ctx, &proto.NextTaskRequest{})
	if err != nil {
		return fmt.Errorf("next task: %w", err)
	}
	for {
		jr, err := c.gc.NextJob(ctx, &proto.NextJobRequest{TaskId: tr.Id})
		if err != nil {
			serr, ok := status.FromError(err)
			// All job has been processed.
			if ok && serr.Code() == codes.NotFound {
				break
			}
			return err
		}

		err = c.pool.Submit(func() {
			StartWorker(c, ToJob(jr))
		})
		if err != nil {
			return fmt.Errorf("submit task: %w", err)
		}
	}
	return nil
}
