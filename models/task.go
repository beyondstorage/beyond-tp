package models

import (
	"context"
	"errors"
	"github.com/dgraph-io/badger/v3"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var TaskDone = errors.New("task done")

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

// Insert will insert task and update all staffs task queue.
func (d *DB) InsertTask(txn *badger.Txn, t *Task) (err error) {
	if txn == nil {
		txn = d.db.NewTransaction(true)
		defer func() {
			err = d.CloseTxn(txn, err)
		}()
	}

	for _, v := range t.StaffIds {
		err = d.InsertStaffTask(txn, v, t.Id)
		if err != nil {
			return
		}
	}

	bs, err := protobuf.Marshal(t)
	if err != nil {
		return err
	}

	if err = txn.Set(TaskKey(t.Id), bs); err != nil {
		return err
	}
	return
}

func (d *DB) UpdateTask(t *Task) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	bs, err := protobuf.Marshal(t)
	if err != nil {
		return err
	}

	// TODO: we need to check task before update it.
	if err = txn.Set(TaskKey(t.Id), bs); err != nil {
		return err
	}
	return txn.Commit()
}

// DeleteTask delete a task by given ID from DB
func (d *DB) DeleteTask(id string) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	// TODO: we will need to check task first.
	if err := txn.Delete(TaskKey(id)); err != nil {
		return err
	}
	return txn.Commit()
}

// GetTask get task from db and parsed into struct with specific ID
func (d *DB) GetTask(id string) (t *Task, err error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(TaskKey(id))
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

func (d *DB) WaitTask(ctx context.Context, taskId string) (err error) {
	_, err = d.GetTask(taskId)
	// If job doesn't exist, we can return directly.
	if err != nil && errors.Is(err, ErrNotFound) {
		return nil
	}
	if err != nil {
		return
	}

	err = d.db.Subscribe(ctx, func(kv *badger.KVList) error {
		for _, v := range kv.Kv {
			if v.Value == nil {
				return TaskDone
			}
		}
		return nil
	}, TaskKey(taskId))
	if err == TaskDone {
		return nil
	}
	return err
}

// This function will be used to elect task leader.
// If there is no leader here, we will use input staff as leader.
// TODO: This logic could be changed.
func (d *DB) ElectTaskLeader(taskId, staffId, staffAddr string) (electedStaffId, electedStaffAddr string, err error) {
	txn := d.db.NewTransaction(true)

	sid, saddr, err := d.getTaskLeader(txn, taskId)
	// We do get a task leader.
	if err == nil {
		return sid, saddr, nil
	}
	// If err is not key not found, other error happened.
	if err != badger.ErrKeyNotFound {
		return "", "", err
	}

	bs, err := protobuf.Marshal(&TaskLeader{
		TaskId:    taskId,
		StaffId:   staffId,
		StaffAddr: staffAddr,
	})
	if err != nil {
		panic("invalid task leader")
	}

	err = txn.Set(TaskLeaderKey(taskId), bs)
	if err != nil {
		return
	}

	err = txn.Commit()
	// Task leader has been updated by other txn, we should discard our changes.
	if err == badger.ErrConflict {
		txn.Discard()
		return d.getTaskLeader(d.db.NewTransaction(false), taskId)
	}
	return staffId, staffAddr, nil
}

func (d *DB) getTaskLeader(txn *badger.Txn, taskId string) (electedStaffId, electedStaffAddr string, err error) {
	item, err := txn.Get(TaskLeaderKey(taskId))
	if err != nil {
		return "", "", err
	}

	tl := &TaskLeader{}

	err = item.Value(func(val []byte) error {
		err = protobuf.Unmarshal(val, tl)
		if err != nil {
			panic("invalid task leader")
		}
		return nil
	})
	if err != nil {
		return
	}

	return tl.StaffId, tl.StaffAddr, nil
}
