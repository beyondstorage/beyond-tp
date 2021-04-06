package models

import (
	"github.com/Xuanwo/go-bufferpool"
)

var pool *bufferpool.Pool

const (
	taskPrefix  = 't'
	staffPrefix = 's'
	jobPrefix   = 'j'
)

var (
	TaskPrefix  = []byte{taskPrefix, ':'}
	StaffPrefix = []byte{staffPrefix, ':'}
	JobPrefix   = []byte{jobPrefix, ':'}
)

func FormatTaskKey(taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(taskPrefix)
	b.AppendByte(':')
	b.AppendString(taskId)

	return b.Bytes()
}

func FormatStaffKey(staffId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(staffPrefix)
	b.AppendByte(':')
	b.AppendString(staffId)

	return b.Bytes()
}

func FormatJobKey(jobId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(jobPrefix)
	b.AppendByte(':')
	b.AppendString(jobId)

	return b.Bytes()
}

func init() {
	// Most key will include a UUID like: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	// Set init size to 64 to prevent alloc extra space.
	pool = bufferpool.New(64)
}
