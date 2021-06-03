package models

import (
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	fs "github.com/beyondstorage/go-service-fs/v3"
	qingstor "github.com/beyondstorage/go-service-qingstor/v3"
	"github.com/beyondstorage/go-storage/v4/types"
)

func FormatStorage(st *Storage) (types.Storager, error) {
	pairs := make([]types.Pair, 0, len(st.Options))
	for _, p := range st.Options {
		pairs = append(pairs, FormatPair(p))
	}

	switch st.Type {
	case StorageType_Fs:
		return fs.NewStorager(pairs...)
	case StorageType_Qingstor:
		return qingstor.NewStorager(pairs...)
	default:
		return nil, errors.New("endpoint type unsupported")
	}
}

func FormatPair(p *Pair) types.Pair {
	return types.Pair{
		Key:   p.Key,
		Value: p.Value,
	}
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

func (x *Pair) MarshalGQL(w io.Writer) {
	graphql.MarshalMap(map[string]interface{}{
		"key":   x.Key,
		"value": x.Value,
	}).MarshalGQL(w)
}

func (x *Pair) UnmarshalGQL(v interface{}) error {
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "key":
			var err error
			x.Key, err = graphql.UnmarshalString(v)
			if err != nil {
				return err
			}
		case "value":
			var err error
			x.Value, err = graphql.UnmarshalString(v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (x *Storage) MarshalGQL(w io.Writer) {
	graphql.MarshalMap(map[string]interface{}{
		"type":    x.Type,
		"options": x.Options,
	}).MarshalGQL(w)
}
