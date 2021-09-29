package task

import "github.com/google/uuid"

const (
	_ = iota
	JobTypeCopyDir
	JobTypeCopySmallFile
	JobTypeCopyLargeFile
	JonTypeCopyPart
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

func ParseCopyPartJob(bs []byte) (t *CopyPartJob) {
	t = &CopyPartJob{}
	MustUnmarshal(bs, t)
	return t
}
