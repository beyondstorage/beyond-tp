package task

import (
	"context"
	"errors"
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
		no, err := nc.Recv()
		if err != nil && errors.Is(err, io.EOF) {
			break
		}
	}
}
