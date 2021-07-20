package models

import (
	"strings"

	"github.com/Xuanwo/go-bufferpool"
)

var pool *bufferpool.Pool

const (
	taskPrefix     = 't'
	staffPrefix    = 's'
	jobPrefix      = 'j'
	identityPrefix = 'i'
)

var (
	TaskPrefix    = []byte{taskPrefix, ':'}
	StaffPrefix   = []byte{staffPrefix, ':'}
	JobPrefix     = []byte{jobPrefix, ':'}
	JobMetaPrefix = []byte{'j', 'm', 't', ':'}
)

// TaskKey Style: t:<task_id>
func TaskKey(taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(taskPrefix)
	b.AppendByte(':')
	b.AppendString(taskId)

	return b.BytesCopy()
}

// TaskLeaderKey Style: t_leader:<task_id>
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

// StaffKey Style: s:<staff_id>
func StaffKey(staffId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(staffPrefix)
	b.AppendByte(':')
	b.AppendString(staffId)

	return b.BytesCopy()
}

// StaffTaskPrefix Style: s_t:<staff_id>:
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

// StaffTaskKey Style: s_t:<staff_id>:<task_id>
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

// JobKey Style: j:<job_id>
func JobKey(jobId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(jobPrefix)
	b.AppendByte(':')
	b.AppendString(jobId)

	return b.BytesCopy()
}

// IdentityKey Style: i:<id_type>:<id_name>
func IdentityKey(idType IdentityType, name string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(identityPrefix)
	b.AppendByte(':')
	b.AppendString(idType.String())
	b.AppendByte(':')
	b.AppendString(name)

	return b.BytesCopy()
}

// IdentityKeyPrefix Style: i:<id_type>:
func IdentityKeyPrefix(idType *IdentityType) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(identityPrefix)
	b.AppendByte(':')

	// append type filter if has
	if idType != nil {
		b.AppendString(idType.String())
		b.AppendByte(':')
	}
	return b.BytesCopy()
}

// JobMetaKey Style: jmt:<job_id>
func JobMetaKey(jobId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendBytes(JobMetaPrefix)
	b.AppendString(jobId)

	return b.BytesCopy()
}

func init() {
	// Most key will include a UUID like: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	// Set init size to 64 to prevent alloc extra space.
	pool = bufferpool.New(64)
}

// GetTaskIDFromStaffTaskKey Style: s_t:<staff_id>:<task_id>
func GetTaskIDFromStaffTaskKey(key string) string {
	results := strings.Split(key, ":")
	if len(results) < 3 {
		return ""
	}
	return results[2]
}
