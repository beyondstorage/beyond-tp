package task

const (
	TypeCopyDir uint32 = iota + 1
	TypeCopyFile
	TypeCopySingleFile
	TypeCopyMultipartFile
	TypeCopyMultipart
)

const (
	JobStatusSucceed uint32 = iota
	JobStatusFailed
)
