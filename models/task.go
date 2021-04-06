package models

import (
	"context"
	"errors"

	"github.com/dgraph-io/badger/v3"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewTask created a task with default value
func NewTask(name string, ty TaskType) *Task {
	now := timestamppb.Now()
	t := Task{
		Id:        uuid.NewString(),
		Name:      name,
		Type:      ty,
		Status:    TaskStatus_Created,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return &t
}

func NewTaskFromBytes(bs []byte) *Task {
	t := &Task{}
	err := protobuf.Unmarshal(bs, t)
	if err != nil {
		panic("invalid task")
	}
	return t
}

// SaveTask save a task into DB
// TODO: should be split into CreateTask and UpdateTask.
func (d *DB) SaveTask(t *Task) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	bs, err := protobuf.Marshal(t)
	if err != nil {
		return err
	}

	if err = txn.Set(FormatTaskKey(t.Id), bs); err != nil {
		return err
	}
	return txn.Commit()
}

// DeleteTask delete a task by given ID from DB
func (d *DB) DeleteTask(id string) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	// TODO: we will need to check task first.
	if err := txn.Delete(FormatTaskKey(id)); err != nil {
		return err
	}
	return txn.Commit()
}

// GetTask get task from db and parsed into struct with specific ID
func (d *DB) GetTask(id string) (t *Task, err error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(FormatTaskKey(id))
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, NewNotFoundErr(id)
		}
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		t = NewTaskFromBytes(val)
		return nil
	})
	return
}

// ListTasks create a db iterator and conduct result tasks
func (d *DB) ListTasks() ([]*Task, error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()
	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	tasks := make([]*Task, 0)

	for it.Seek(TaskPrefix); it.ValidForPrefix(TaskPrefix); it.Next() {
		item := it.Item()
		err := item.Value(func(v []byte) error {
			tasks = append(tasks, NewTaskFromBytes(v))
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func (d *DB) SubscribeTask(ctx context.Context, fn func(t *Task)) (err error) {
	return d.db.Subscribe(ctx, func(kv *badger.KVList) error {
		for _, v := range kv.Kv {
			fn(NewTaskFromBytes(v.Value))
		}
		return nil
	}, TaskPrefix)
}

// This function will be used to elect task leader.
// If there is no leader here, we will use input staff as leader.
// TODO: This logic could be changed.
func (d *DB) ElectTaskLeader(taskId, staffId, staffAddr string) (electedStaffId, electedStaffAddr string, err error) {
	// Require a lock
	panic("implement me")
}
