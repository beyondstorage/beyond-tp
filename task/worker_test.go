package task

import (
	"context"
	"github.com/aos-dev/dm/models"
	"github.com/aos-dev/go-toolbox/zapcontext"
	"github.com/google/uuid"
	"testing"
)

func setupPortal(t *testing.T) *Manager {
	p, err := NewManager(context.Background(), ManagerConfig{
		Host:     "localhost",
		GrpcPort: 7000,
	})
	if err != nil {
		t.Error(err)
	}

	go p.Serve(context.Background())
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
		w.Start(ctx)
	}

	copyFileTask := &models.Task{
		Id:     uuid.NewString(),
		Type:   models.TaskType_CopyDir,
		Status: models.TaskStatus_Ready,
		Storages: []*models.Storage{
			{Type: models.StorageType_Fs, Options: []*models.Pair{{Key: "work_dir", Value: "/tmp/b/"}}},
			{Type: models.StorageType_Fs, Options: []*models.Pair{{Key: "work_dir", Value: "/tmp/c/"}}},
		},
	}

	p.db.SaveTask(copyFileTask)
}
