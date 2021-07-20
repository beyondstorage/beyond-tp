package task

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/beyondstorage/go-toolbox/zapcontext"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

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
			t.Error(err)
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
		t.Errorf("insert task: %v", err)
	}

	task.UpdatedAt = timestamppb.Now()
	task.Status = models.TaskStatus_Running

	txn := p.db.NewTxn(true)
	err = p.db.UpdateTask(txn, task)
	if err != nil {
		t.Error("save task", zap.String("id", task.Id), zap.Error(err))
		txn.Discard()
		t.Fatal(err)
	}

	time.Sleep(time.Second)
	for _, staffId := range task.StaffIds {
		err = p.db.InsertStaffTask(txn, staffId, task.Id)
		if err != nil {
			t.Error("insert staff task",
				zap.String("task", task.Id), zap.String("staff", staffId), zap.Error(err))
			txn.Discard()
			t.Fatal(err)
		}
	}
	txn.Commit()

	err = p.db.WaitTask(ctx, task.Id)
	if err != nil {
		t.Errorf("wait task: %v", err)
	}
	t.Logf("task has been finished")

	err = p.Stop(ctx)
	if err != nil {
		t.Errorf("stop: %v", err)
	}
}
