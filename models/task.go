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

	"github.com/aos-dev/dm/proto"
	"github.com/aos-dev/dm/task"
	"github.com/dgraph-io/badger/v3"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

const (
	StatusUnknown = iota
	StatusCreated
	StatusRunning
	StatusFinished
	StatusStopped
	StatusError
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
	case StatusError:
		return "error"
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
	case "error":
		res = StatusError
	default:
		res = StatusUnknown
	}
	*ts = res
}

// IsRunning assert whether task status is running
func (ts *TaskStatus) IsRunning() bool {
	if ts == nil {
		return false
	}
	return *ts == StatusRunning
}

// Task contains info of data migration task
type Task struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Type      TaskType               `json:"type"`
	Status    TaskStatus             `json:"status"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Src       Endpoint               `json:"src"`
	Dst       Endpoint               `json:"dst"`
	Options   map[string]interface{} `json:"options,omitempty"`
}

// Endpoint contains info to create an endpoint
type Endpoint struct {
	Type    ServiceType `json:"type"`
	Options interface{} `json:"options,omitempty"`
}

// parse Endpoint into *proto.Endpoint
func (e Endpoint) parse() *proto.Endpoint {
	// ensure handle e.Option as map[string]interface{}
	opt := make(map[string]interface{})
	if e.Options != nil {
		opt = e.Options.(map[string]interface{})
	}

	pairs := make([]*proto.Pair, 0, len(opt)+1) // +1 for work dir inject

	// conduct pairs with endpoint's options
	for k, v := range opt {
		pairs = append(pairs, &proto.Pair{Key: k, Value: v.(string)})
	}

	return &proto.Endpoint{Type: e.Type.String(), Pairs: pairs}
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
		Src:       Endpoint{Options: make(map[string]interface{})},
		Dst:       Endpoint{Options: make(map[string]interface{})},
		Options:   make(map[string]interface{}),
	}
	return &t
}

// FormatProtoTask conduct task into *proto.Task
func (t Task) FormatProtoTask() (*proto.Task, error) {
	// TODO: conduct other tasks, such as move or sync
	copyFileJob := &proto.CopyDir{
		Src:       0,
		Dst:       1,
		SrcPath:   "",
		DstPath:   "",
		Recursive: t.Options["recursive"].(bool),
	}
	content, err := protobuf.Marshal(copyFileJob)
	if err != nil {
		return nil, err
	}

	copyFileTask := &proto.Task{
		Id: uuid.NewString(),
		Endpoints: []*proto.Endpoint{
			t.Src.parse(),
			t.Dst.parse(),
		},
		Job: &proto.Job{
			Id:      uuid.NewString(),
			Type:    uint32(t.Type),
			Content: content,
		},
	}
	return copyFileTask, nil
}

// SaveTask save a task into DB
func (d *DB) SaveTask(t *Task) error {
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
			t := &Task{}
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
			tasks = append(tasks, t)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
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

type ServiceType string

// String implement Stringer for ServiceType
func (st ServiceType) String() string {
	return string(st)
}

// Parse type string into ServiceType
func (st *ServiceType) Parse(t string) {
	*st = ServiceType(t)
}

func (st ServiceType) MarshalGQL(w io.Writer) {
	_, err := w.Write([]byte(strconv.Quote(st.String())))
	// handle error as panic
	if err != nil {
		panic(err)
	}
}

func (st *ServiceType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		st.Parse(strings.ToLower(v))
		return nil
	case ServiceType:
		*st = v
		return nil
	default:
		return fmt.Errorf("%T is not a string", v)
	}
}
