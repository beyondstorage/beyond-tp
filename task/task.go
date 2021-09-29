package task

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	_ = iota
	TaskTypeCopyDir
)

const (
	_ = iota
	TaskStatusReady
)

type Task struct {
	Id                 string
	Name               string
	Type               int
	Status             int
	CreatedAt          time.Time `msgpack:"cat"`
	UpdatedAt          time.Time `msgpack:"uat"`
	StorageConnections []string  `msgpack:"sc"`

	// msgpack trick to omit all empty fields in struct.
	_msgpack struct{} `msgpack:",omitempty"`
}

func NewTask(name string, typ int) *Task {
	now := time.Now()

	return &Task{
		Id:        uuid.NewString(),
		Name:      name,
		Type:      typ,
		Status:    TaskStatusReady,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func FormatTask(t *Task) []byte {
	bs, err := msgpack.Marshal(t)
	if err != nil {
		panic(fmt.Errorf("marshal task: %w", err))
	}
	return bs
}

func ParseTask(bs []byte) (t *Task) {
	t = &Task{}
	err := msgpack.Unmarshal(bs, t)
	if err != nil {
		panic(fmt.Errorf("unmarshal task: %w", err))
	}
	return t
}
