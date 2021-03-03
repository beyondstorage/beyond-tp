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

	"github.com/dgraph-io/badger/v3"
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
	w.Write([]byte(strconv.Quote(ts.String())))
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

// SaveTask save a task into DB
func (h *DBHandler) SaveTask(t Task) error {
	return h.db.Update(func(txn *badger.Txn) error {
		res, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return txn.Set(t.FormatKey(), res)
	})
}

// DeleteTask delete a task from DB
func (h *DBHandler) DeleteTask(t *Task) error {
	txn := h.db.NewTransaction(true)
	defer txn.Discard()
	// try to get task first
	item, err := txn.Get(t.FormatKey())
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return fmt.Errorf("%s: %w", err.Error(), ErrNotFound)
		}
		return err
	}
	if err = item.Value(func(val []byte) error {
		return json.Unmarshal(val, t)
	}); err != nil {
		return err
	}
	// then delete
	if err = txn.Delete(t.FormatKey()); err != nil {
		return err
	}
	return txn.Commit()
}

// SetStatus transfer int into TaskStatus and set it for task
func (t *Task) SetStatus(status int) {
	t.Status = TaskStatus(status)
}

// GetTask get task from db and parsed into struct with specific ID
func (h *DBHandler) GetTask(id string) (*Task, error) {
	t := &Task{ID: id}
	txn := h.db.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get(t.FormatKey())
	if err != nil {
		// handle not found error manually
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, fmt.Errorf("%s: %w", err.Error(), ErrNotFound)
		}
		return nil, err
	}
	err = item.Value(func(val []byte) error {
		return json.Unmarshal(val, t)
	})
	return t, err
}

// ListTasks create a db iterator and conduct result tasks
func (h *DBHandler) ListTasks() ([]*Task, error) {
	txn := h.db.NewTransaction(false)
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
