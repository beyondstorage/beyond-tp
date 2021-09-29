package task

import (
	"fmt"
	"github.com/beyondstorage/beyond-tp/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	TaskTypeInvalid = iota
	TaskTypeCopyDir
)

const (
	TaskStatusInvalid = iota
	TaskStatusReady
	TaskStatusRunning
)

type Task struct {
	Id        string
	Name      string
	Type      int
	Status    int
	CreatedAt time.Time `msgpack:"cat"`
	UpdatedAt time.Time `msgpack:"uat"`
	Storages  []string  `msgpack:"ss"`

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

func FromTask(t *Task) *proto.TaskReply {
	return &proto.TaskReply{
		Id:        t.Id,
		Name:      t.Name,
		Type:      int64(t.Type),
		Status:    int64(t.Status),
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
		Storages:  t.Storages,
	}
}

func ToTask(tr *proto.TaskReply) *Task {
	return &Task{
		Id:        tr.Id,
		Name:      tr.Name,
		Type:      int(tr.Type),
		Status:    int(tr.Status),
		CreatedAt: tr.CreatedAt.AsTime(),
		UpdatedAt: tr.UpdatedAt.AsTime(),
		Storages:  tr.Storages,
	}
}
