package task

import "github.com/Xuanwo/go-bufferpool"

// Most key will include a UUID like: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
// Set init size to 64 to prevent alloc extra space.
var pool = bufferpool.New(64)

const (
	separator = ':'

	prefixTask = 't'
	prefixJob  = 'j'
	prefixMeta = 'm'
)

var (
	PrefixTask = []byte{prefixTask, separator}
)

// KeyTask format: t:<task-id>
func KeyTask(taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendBytes(PrefixTask)
	b.AppendString(taskId)

	return b.BytesCopy()
}

// PrefixJob format: j:<task-id>:
func PrefixJob(taskId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(prefixJob)
	b.AppendByte(separator)
	b.AppendString(taskId)
	b.AppendByte(separator)

	return b.BytesCopy()
}

// KeyJob format: j:<task-id>:<job-id>
func KeyJob(taskId, jobId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendBytes(PrefixJob(taskId))
	b.AppendString(jobId)

	return b.BytesCopy()
}

// PrefixMeta format: m:<task-id>:<job-id>:
func PrefixMeta(taskId, jobId string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendByte(prefixMeta)
	b.AppendByte(separator)
	b.AppendString(taskId)
	b.AppendByte(separator)
	b.AppendString(jobId)
	b.AppendByte(separator)

	return b.BytesCopy()
}

// KeyMeta format: j:<task-id>:<job-id>:<meta-key>
func KeyMeta(taskId, jobId, metaKey string) []byte {
	b := pool.Get()
	defer b.Free()

	b.AppendBytes(PrefixMeta(taskId, jobId))
	b.AppendString(metaKey)

	return b.BytesCopy()
}
