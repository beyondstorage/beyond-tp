package models

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	_ "github.com/beyondstorage/go-service-azblob/v2"
	_ "github.com/beyondstorage/go-service-cos/v2"
	_ "github.com/beyondstorage/go-service-dropbox/v2"
	_ "github.com/beyondstorage/go-service-fs/v3"
	_ "github.com/beyondstorage/go-service-gcs/v2"
	_ "github.com/beyondstorage/go-service-kodo/v2"
	_ "github.com/beyondstorage/go-service-oss/v2"
	_ "github.com/beyondstorage/go-service-qingstor/v3"
	_ "github.com/beyondstorage/go-service-s3/v2"
	"github.com/beyondstorage/go-storage/v4/services"
	"github.com/beyondstorage/go-storage/v4/types"
)

func FormatStorage(st *Storage) (types.Storager, error) {
	store, err := services.NewStoragerFromString(st.Connection)
	if err != nil {
		return nil, fmt.Errorf("init storage: %v", err)
	}
	return store, nil
}

func (x TaskType) MarshalGQL(w io.Writer) {
	if x == TaskType_InvalidTaskType {
		panic("invalid task type")
	}

	_, _ = w.Write([]byte(x.String()))
}

func (x *TaskType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*x = TaskType(TaskType_value[v])
		return nil
	default:
		return fmt.Errorf("%T is not a string", v)
	}
}

func (x TaskStatus) MarshalGQL(w io.Writer) {
	if x == TaskStatus_InvalidTaskStatus {
		panic("invalid task type")
	}

	_, _ = w.Write([]byte(x.String()))
}

func (x *TaskStatus) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*x = TaskStatus(TaskStatus_value[v])
		return nil
	default:
		return fmt.Errorf("%T is not a string", v)
	}
}

func (x StorageType) MarshalGQL(w io.Writer) {
	if x == StorageType_InvalidStorageType {
		panic("invalid task type")
	}

	_, _ = w.Write([]byte(x.String()))
}

func (x *StorageType) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*x = StorageType(StorageType_value[v])
		return nil
	default:
		return fmt.Errorf("%T is not a string", v)
	}
}

func (x *Storage) MarshalGQL(w io.Writer) {
	graphql.MarshalMap(map[string]interface{}{
		"type": x.Type,
	}).MarshalGQL(w)
}
