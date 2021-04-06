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

// Style: t:<task_id>
func TaskKey(taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(taskPrefix)
	b.AppendByte(':')
	b.AppendString(taskId)

	return b.BytesCopy()
}

// Style: t_leader:<task_id>
func TaskLeaderKey(taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(taskPrefix)
	b.AppendByte('_')
	b.AppendString("leader")
	b.AppendByte(':')
	b.AppendString(taskId)

	return b.BytesCopy()
}

// Style: s:<staff_id>
func StaffKey(staffId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(staffPrefix)
	b.AppendByte(':')
	b.AppendString(staffId)

	return b.BytesCopy()
}

// Style: s_t:<staff_id>:<task_id>
func StaffTaskPrefix(staffId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(staffPrefix)
	b.AppendByte('_')
	b.AppendByte(taskPrefix)
	b.AppendByte(':')
	b.AppendString(staffId)
	b.AppendByte(':')

	return b.BytesCopy()
}

// Style: st:<staff_id>:<task_id>
func StaffTaskKey(staffId, taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(staffPrefix)
	b.AppendByte('_')
	b.AppendByte(taskPrefix)
	b.AppendByte(':')
	b.AppendString(staffId)
	b.AppendByte(':')
	b.AppendString(taskId)

	return b.BytesCopy()
}

func JobKey(jobId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(jobPrefix)
	b.AppendByte(':')
	b.AppendString(jobId)

	return b.BytesCopy()
}

func init() {
	// Most key will include a UUID like: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	// Set init size to 64 to prevent alloc extra space.
	pool = bufferpool.New(64)
}
