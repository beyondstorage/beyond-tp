package rpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/beyondstorage/beyond-tp/proto"
	"github.com/beyondstorage/beyond-tp/task"
)

func FromTask(t *task.Task) *proto.TaskReply {
	return &proto.TaskReply{
		Id:                 t.Id,
		Name:               t.Name,
		Type:               int64(t.Type),
		Status:             int64(t.Status),
		CreatedAt:          timestamppb.New(t.CreatedAt),
		UpdatedAt:          timestamppb.New(t.UpdatedAt),
		StorageConnections: t.StorageConnections,
	}
}

func ToTask(tr *proto.TaskReply) *task.Task {
	return &task.Task{
		Id:                 tr.Id,
		Name:               tr.Name,
		Type:               int(tr.Type),
		Status:             int(tr.Status),
		CreatedAt:          tr.CreatedAt.AsTime(),
		UpdatedAt:          tr.UpdatedAt.AsTime(),
		StorageConnections: tr.StorageConnections,
	}
}

func FromJob(j *task.Job) *proto.JobReply {
	return &proto.JobReply{
		Id:      j.Id,
		TaskId:  j.TaskId,
		Type:    int64(j.Type),
		Content: j.Content,
	}
}

func ToJob(jr *proto.JobReply) *task.Job {
	return &task.Job{
		Id:      jr.Id,
		TaskId:  jr.TaskId,
		Type:    int(jr.Type),
		Content: jr.Content,
	}
}
