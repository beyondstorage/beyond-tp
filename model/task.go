package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	StatusBroken
)

const taskPrefix = "t"

type User struct {
	ID   int
	Name string
	Addr string
}

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    int       `json:"status"`
	StatusStr string    `json:"status_str"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tasks []Task

// FormatKey format db key for task
func (t Task) FormatKey() []byte {
	b := new(bytes.Buffer)
	b.WriteString(taskPrefix)
	b.WriteString(":")
	b.WriteString(t.ID)
	return b.Bytes()
}

// NumToStatus check status num and transfer to status
func (t Task) NumToStatus() string {
	switch t.Status {
	case StatusCreated:
		return "created"
	case StatusRunning:
		return "running"
	case StatusFinished:
		return "finished"
	case StatusStopped:
		return "stopped"
	case StatusBroken:
		return "broken"
	default:
		return "unknown"
	}
}

// ParseStatus parse status string into num
func (t Task) ParseStatus() int {
	status := strings.ToLower(t.StatusStr)
	switch status {
	case "created":
		return StatusCreated
	case "running":
		return StatusRunning
	case "finished":
		return StatusFinished
	case "stopped":
		return StatusStopped
	case "broken":
		return StatusBroken
	default:
		return StatusUnknown
	}
}

func (t Task) Save() error {
	return DB.Update(func(txn *badger.Txn) error {
		res, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return txn.Set(t.FormatKey(), res)
	})
}

func (t Task) Delete() error {
	return DB.Update(func(txn *badger.Txn) error {
		return txn.Delete(t.FormatKey())
	})
}

// Get method get task from db and parsed into struct
func (t *Task) Get() error {
	if t.ID == "" {
		return fmt.Errorf("id required: %w", ErrKeyRequired)
	}
	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(t.FormatKey())
		if err != nil {
			// handle not found error manually
			if errors.Is(err, badger.ErrKeyNotFound) {
				return fmt.Errorf("%s: %w", err.Error(), ErrNotFound)
			}
			return err
		}
		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, t)
		})
		return err
	})
	return err
}

// GetTaskByID call task.Get
func GetTaskByID(id string) (Task, error) {
	task := Task{ID: id}
	err := task.Get()
	return task, err
}

// GetTaskList create a db iterator and conduct result tasks
func GetTaskList() ([]Task, error) {
	tasks := make([]Task, 0)
	err := DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(taskPrefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				task := Task{}
				err := json.Unmarshal(v, &task)
				if err != nil {
					return err
				}
				tasks = append(tasks, task)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
