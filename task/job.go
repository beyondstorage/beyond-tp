package task

import (
	"github.com/google/uuid"

	"github.com/beyondstorage/beyond-tp/proto"
)

const (
	_ = iota
	JobTypeCopyDir
	JobTypeCopySmallFile
	JobTypeCopyLargeFile
	JobTypeCopyPart
)

type Job struct {
	Id      string
	TaskId  string
	Type    int
	Content []byte
}

func NewJob(taskId string, typ int, content []byte) *Job {
	return &Job{
		Id:      uuid.NewString(),
		TaskId:  taskId,
		Type:    typ,
		Content: content,
	}
}

func FormatJob(t *Job) []byte {
	return MustMarshal(t)
}

func ParseJob(bs []byte) (t *Job) {
	t = &Job{}
	MustUnmarshal(bs, t)
	return t
}

func ToProtoJob(j *Job) *proto.Job {
	return &proto.Job{
		Id:      j.Id,
		TaskId:  j.TaskId,
		Type:    int64(j.Type),
		Content: j.Content,
	}
}

func FromProtoJob(jr *proto.Job) *Job {
	return &Job{
		Id:      jr.Id,
		TaskId:  jr.TaskId,
		Type:    int(jr.Type),
		Content: jr.Content,
	}
}

type CopyDirJob struct {
	SrcPath string
	DstPath string
}

type CopySmallFileJob struct {
	SrcPath string
	DstPath string
	Size    int64
}

type CopyLargeFileJob struct {
	SrcPath string
	DstPath string
	Size    int64
}

type CopyPartJob struct {
	SrcPath     string
	DstPath     string
	MultipartId string
	Size        int64
	Index       int
	Offset      int64
}
