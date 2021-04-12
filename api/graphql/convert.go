package graphql

import (
	"fmt"

	"github.com/aos-dev/dm/models"
)

func parseTaskType(tt TaskType) models.TaskType {
	switch tt {
	case TaskTypeCopyDir:
		return models.TaskType_CopyDir
	default:
		return models.TaskType_InvalidTaskType
	}
}

func formatTaskType(tt models.TaskType) TaskType {
	switch tt {
	case models.TaskType_CopyDir:
		return TaskTypeCopyDir
	default:
		panic(fmt.Errorf("task type %s is invalid", tt.String()))
	}
}

func formatTaskStatus(ts models.TaskStatus) TaskStatus {
	switch ts {
	case models.TaskStatus_Created:
		return TaskStatusCreated
	case models.TaskStatus_Ready:
		return TaskStatusReady
	case models.TaskStatus_Running:
		return TaskStatusRunning
	case models.TaskStatus_Finished:
		return TaskStatusFinished
	case models.TaskStatus_Stopped:
		return TaskStatusStopped
	case models.TaskStatus_Error:
		return TaskStatusError
	default:
		panic(fmt.Errorf("task status %s is invalid", ts.String()))
	}
}

func parseStorageType(st StorageType) models.StorageType {
	switch st {
	case StorageTypeFs:
		return models.StorageType_Fs
	case StorageTypeQingstor:
		return models.StorageType_Qingstor
	default:
		return models.StorageType_InvalidStorageType
	}
}

func formatStorageType(st models.StorageType) StorageType {
	switch st {
	case models.StorageType_Fs:
		return StorageTypeFs
	case models.StorageType_Qingstor:
		return StorageTypeQingstor
	default:
		panic(fmt.Errorf("storage type %s is invalid", st.String()))
	}
}

func parsePairInput(pi *PairInput) *models.Pair {
	return &models.Pair{
		Key:   pi.Key,
		Value: pi.Value,
	}
}

func formatPair(pi *models.Pair) *Pair {
	return &Pair{
		Key:   pi.Key,
		Value: pi.Value,
	}
}

func parsePairsInput(pis []*PairInput) []*models.Pair {
	ps := make([]*models.Pair, 0, len(pis))
	for _, v := range pis {
		ps = append(ps, parsePairInput(v))
	}
	return ps
}

func formatPairs(pis []*models.Pair) []*Pair {
	ps := make([]*Pair, 0, len(pis))
	for _, v := range pis {
		ps = append(ps, formatPair(v))
	}
	return ps
}

func parseStorageInput(si *StorageInput) *models.Storage {
	return &models.Storage{
		Type:    parseStorageType(si.Type),
		Options: parsePairsInput(si.Options),
	}
}

func formatStorage(ms *models.Storage) *Storage {
	return &Storage{
		Type:    formatStorageType(ms.Type),
		Options: formatPairs(ms.Options),
	}
}

func parseStoragesInput(si []*StorageInput) []*models.Storage {
	ss := make([]*models.Storage, 0, len(si))
	for _, v := range si {
		ss = append(ss, parseStorageInput(v))
	}
	return ss
}

func formatStorages(si []*models.Storage) []*Storage {
	ss := make([]*Storage, 0, len(si))
	for _, v := range si {
		ss = append(ss, formatStorage(v))
	}
	return ss
}

func formatTask(t *models.Task) *Task {
	return &Task{
		ID:        t.Id,
		Name:      t.Name,
		Type:      formatTaskType(t.Type),
		Status:    formatTaskStatus(t.Status),
		CreatedAt: t.CreatedAt.AsTime(),
		UpdatedAt: t.UpdatedAt.AsTime(),
		Storages:  formatStorages(t.Storages),
		Options:   formatPairs(t.Options),
	}
}

func formatTasks(t []*models.Task) []*Task {
	ts := make([]*Task, 0, len(t))
	for _, v := range t {
		ts = append(ts, formatTask(v))
	}
	return ts
}
