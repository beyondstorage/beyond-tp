package task

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/beyondstorage/beyond-tp/models"
)

func setupManager(t *testing.T) *Manager {
	os.RemoveAll("/tmp/badger")

	p, err := NewManager(context.Background(), ManagerConfig{
		Host:         "localhost",
		GrpcPort:     7000,
		DatabasePath: "/tmp/badger",
	})
	if err != nil {
		t.Error(err)
	}

	return p
}

// This is not a really unit test, just for developing, SHOULD be removed.
func TestWorker(t *testing.T) {
	zapcontext.SetFactoryFunction(func() *zap.Logger {
		logger, _ := zap.NewDevelopment()
		return logger
	})

	p := setupManager(t)

	ctx := context.Background()
	_ = zapcontext.From(ctx)

	staffIds := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		w, err := NewStaff(ctx, StaffConfig{
			Host:        "localhost",
			ManagerAddr: "localhost:7000",
			DataPath:    "/tmp/badger",
		})
		if err != nil {
			t.Fatal(err)
		}

		staffIds = append(staffIds, w.id)

		go w.Start(ctx)
	}

	task := &models.Task{
		Id:       uuid.NewString(),
		Type:     models.TaskType_CopyDir,
		Status:   models.TaskStatus_Ready,
		StaffIds: staffIds,
		Storages: []*models.Storage{
			{Type: models.StorageType_Fs, Options: []*models.Pair{{Key: "work_dir", Value: "/tmp/b/"}}},
			{Type: models.StorageType_Fs, Options: []*models.Pair{{Key: "work_dir", Value: "/tmp/c/"}}},
		},
	}

	err := p.db.InsertTask(nil, task)
	if err != nil {
		t.Fatalf("insert task: %v", err)
	}

	time.Sleep(time.Second)

	err = p.db.RunTask(task.Id)
	if err != nil {
		t.Fatalf("run task: %v", err)
	}

	err = p.db.WaitTask(ctx, task.Id)
	if err != nil {
		t.Fatalf("wait task: %v", err)
	}
	t.Logf("task has been finished")

	err = p.Stop(ctx)
	if err != nil {
		t.Fatalf("stop: %v", err)
	}
}
