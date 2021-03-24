package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/aos-dev/noah/task"
	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
)

const (
	StatusUnknown = iota
	StatusCreated
	StatusRunning
	StatusFinished
	StatusStopped
)

const taskPrefix = "t"

// TaskStatus represent status of task
type TaskStatus int

func (ts TaskStatus) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(strconv.Quote(ts.String())))
	// handle error as panic
	if err != nil {
		panic(err)
	}
}

func (ts *TaskStatus) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		ts.Parse(strings.ToLower(v))
		return nil
	case int:
		*ts = TaskStatus(v)
		return nil
	case TaskStatus:
		*ts = v
		return nil
	default:
		return fmt.Errorf("%T is not a int or string", v)
	}
}

// String implement Stringer for TaskStatus
func (ts TaskStatus) String() string {
	switch ts {
	case StatusCreated:
		return "created"
	case StatusRunning:
		return "running"
	case StatusFinished:
		return "finished"
	case StatusStopped:
		return "stopped"
	default:
		return "unknown"
	}
}

// Parse status string into TaskStatus
func (ts *TaskStatus) Parse(status string) {
	var res TaskStatus
	switch strings.ToLower(status) {
	case "created":
		res = StatusCreated
	case "running":
		res = StatusRunning
	case "finished":
		res = StatusFinished
	case "stopped":
		res = StatusStopped
	default:
		res = StatusUnknown
	}
	*ts = res
}

// Task contains info of data migration task
type Task struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// FormatKey format db key for task
func (t Task) FormatKey() []byte {
	b := new(bytes.Buffer)
	b.WriteString(taskPrefix)
	b.WriteString(":")
	b.WriteString(t.ID)
	return b.Bytes()
}

// NewTask created a task with default value
func NewTask() *Task {
	now := time.Now()
	t := Task{
		ID:        uuid.NewString(),
		Status:    StatusCreated, // set StatusCreated as default value
		CreatedAt: now,
		UpdatedAt: now,
	}
	return &t
}

// CreateTask save a task into DB
func (d *DB) CreateTask(t *Task) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	res, err := json.Marshal(t)
	if err != nil {
		return err
	}

	if err = txn.Set(t.FormatKey(), res); err != nil {
		return err
	}
	return txn.Commit()
}

// DeleteTask delete a task by given ID from DB
func (d *DB) DeleteTask(id string) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()
	// try to get task first
	t := Task{ID: id}
	if err := txn.Delete(t.FormatKey()); err != nil {
		return err
	}
	return txn.Commit()
}

// GetTask get task from db and parsed into struct with specific ID
func (d *DB) GetTask(id string) (*Task, error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	t := &Task{ID: id}
	item, err := txn.Get(t.FormatKey())
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, NewNotFoundErr(id)
		}
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		return json.Unmarshal(val, t)
	})
	return t, err
}

// ListTasks create a db iterator and conduct result tasks
func (d *DB) ListTasks() ([]*Task, error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()
	it := txn.NewIterator(badger.DefaultIteratorOptions)
	defer it.Close()

	tasks := make([]*Task, 0)
	prefix := []byte(taskPrefix)
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		item := it.Item()
		err := item.Value(func(v []byte) error {
			task := &Task{}
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			tasks = append(tasks, task)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func (d *DB) SetTaskStatus(id string, status TaskStatus) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	t := Task{ID: id}
	item, err := txn.Get(t.FormatKey())
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return NewNotFoundErr(id)
		}
		return err
	}
	err = item.Value(func(val []byte) error {
		return json.Unmarshal(val, &t)
	})
	if err != nil {
		return err
	}

	t.Status = status

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = txn.Set(t.FormatKey(), data)
	if err != nil {
		return err
	}
	return nil
}

type TaskType uint32

// String implement Stringer for TaskType
func (tt TaskType) String() string {
	switch uint32(tt) {
	case task.TypeCopyDir:
		return "copy_dir"
	case task.TypeCopyFile:
		return "copy_file"
	default:
		return "unknown"
	}
}

// Parse type string into TaskType
func (tt *TaskType) Parse(t string) {
	var res uint32
	switch strings.ToLower(t) {
	case "copy_file":
		res = task.TypeCopyFile
	default: // copy dir as default
		res = task.TypeCopyDir
	}
	*tt = TaskType(res)
}

func (tt TaskType) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(strconv.Quote(tt.String())))
	// handle error as panic
	if err != nil {
		panic(err)
	}
}

func (tt *TaskType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		tt.Parse(strings.ToLower(v))
		return nil
	case uint32:
		*tt = TaskType(v)
		return nil
	case TaskType:
		*tt = v
		return nil
	default:
		return fmt.Errorf("%T is not a uint32 or string", v)
	}
}
