package task

import (
	"context"
	"fmt"
	"github.com/beyondstorage/beyond-tp/proto"
	"io"
)

type Client struct {
	gc proto.AgentClient // gRPC Client
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
	return nil
}
