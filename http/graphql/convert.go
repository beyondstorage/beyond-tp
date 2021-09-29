package graphql

import (
	"fmt"

	"github.com/beyondstorage/beyond-tp/task"
)

func parseTaskType(tt TaskType) int {
	switch tt {
	case TaskTypeCopyDir:
		return task.TaskTypeCopyDir
	default:
		return task.TaskTypeInvalid
	}
}

func formatTaskType(tt int) TaskType {
	switch tt {
	case task.TaskTypeCopyDir:
		return TaskTypeCopyDir
	default:
		panic(fmt.Errorf("task type %d is invalid", tt))
	}
}

func formatTaskStatus(ts int) TaskStatus {
	switch ts {
	case task.TaskStatusReady:
		return TaskStatusReady
	case task.TaskStatusRunning:
		return TaskStatusRunning
	default:
		panic(fmt.Errorf("task status %d is invalid", ts))
	}
}

func formatTask(t *task.Task) *Task {
	return &Task{
		ID:        t.Id,
		Name:      t.Name,
		Type:      formatTaskType(t.Type),
		Status:    formatTaskStatus(t.Status),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Storages:  t.Storages,
	}
}
