package task

import (
	"context"
	"github.com/beyondstorage/beyond-tp/proto"
	"github.com/beyondstorage/go-storage/v4/pairs"
	"github.com/beyondstorage/go-storage/v4/types"
	"io"
	"sync"
)

type Runner struct {
	c *Client

	src types.Storager
	dst types.Storager

	wg *sync.WaitGroup
}

func (rn *Runner) Sync()  {}
func (rn *Runner) Async() {}
func (rn *Runner) Await() {}

func (rn *Runner) HandleCopyPartJob(ctx context.Context, j *Job) (err error) {
	ij := ParseCopyPartJob(j.Content)

	mu := rn.dst.(types.Multiparter)

	pr, pw := io.Pipe()

	go func() {
		_, err := rn.src.Read(ij.SrcPath, pw, pairs.WithSize(ij.Size), pairs.WithOffset(ij.Offset))
		if err != nil {
			return
		}
	}()

	o := rn.dst.Create(ij.DstPath, pairs.WithMultipartID(ij.MultipartId))

	_, part, err := mu.WriteMultipart(o, pr, ij.Size, ij.Index)
	if err != nil {
		return
	}

	_, err = rn.c.gc.SetMeta(ctx, &proto.MetaEntry{
		Key:   nil,
		Value: []byte(part.ETag),
	})
	if err != nil {
		return
	}
	return nil
}
