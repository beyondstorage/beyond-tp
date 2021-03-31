package task

import (
	"context"
	"github.com/aos-dev/dm/proto"
	"github.com/aos-dev/go-toolbox/zapcontext"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"testing"
)

func setupPortal(t *testing.T) *Manager {
	p, err := NewManager(context.Background(), ManagerConfig{
		Host:      "localhost",
		GrpcPort:  7000,
		QueuePort: 7010,
	})
	if err != nil {
		t.Error(err)
	}

	return p
}

// This is not a really unit test, just for developing, SHOULD be removed.
func TestWorker(t *testing.T) {
	p := setupPortal(t)

	ctx := context.Background()
	_ = zapcontext.From(ctx)

	for i := 0; i < 3; i++ {
		w, err := NewStaff(ctx, StaffConfig{
			Host:        "localhost",
			ManagerAddr: "localhost:7000",
		})
		if err != nil {
			t.Error(err)
		}
		err = w.Connect(ctx)
		if err != nil {
			t.Error(err)
		}
	}

	copyFileJob := &proto.CopyDir{
		Src:       0,
		Dst:       1,
		SrcPath:   "",
		DstPath:   "",
		Recursive: true,
	}
	content, err := protobuf.Marshal(copyFileJob)
	if err != nil {
		t.Error(err)
	}

	copyFileTask := &proto.Task{
		Id: uuid.NewString(),
		Endpoints: []*proto.Endpoint{
			{Type: "fs", Pairs: []*proto.Pair{{Key: "work_dir", Value: "/tmp/b/"}}},
			{Type: "fs", Pairs: []*proto.Pair{{Key: "work_dir", Value: "/tmp/c/"}}},
		},
		Job: &proto.Job{
			Id:      uuid.NewString(),
			Type:    TypeCopyDir,
			Content: content,
		},
	}

	t.Log("Start publish")
	err = p.Publish(ctx, copyFileTask)
	if err != nil {
		t.Error(err)
	}

	err = p.Wait(ctx, copyFileTask)
	if err != nil {
		t.Error(err)
	}
}
