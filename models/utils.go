package models

import (
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	azblob "github.com/beyondstorage/go-service-azblob/v2"
	cos "github.com/beyondstorage/go-service-cos/v2"
	dropbox "github.com/beyondstorage/go-service-dropbox/v2"
	fs "github.com/beyondstorage/go-service-fs/v3"
	gcs "github.com/beyondstorage/go-service-gcs/v2"
	kodo "github.com/beyondstorage/go-service-kodo/v2"
	oss "github.com/beyondstorage/go-service-oss/v2"
	qingstor "github.com/beyondstorage/go-service-qingstor/v3"
	s3 "github.com/beyondstorage/go-service-s3/v2"
	"github.com/beyondstorage/go-storage/v4/types"
)

func FormatStorage(st *Storage) (types.Storager, error) {
	pairs := make([]types.Pair, 0, len(st.Options))
	for _, p := range st.Options {
		pairs = append(pairs, FormatPair(p))
	}

	switch st.Type {
	case StorageType_Azblob:
		pairs = append(pairs, azblob.WithStorageFeatures(azblob.StorageFeatures{
			VirtualDir: true,
		}))
		return azblob.NewStorager(pairs...)
	case StorageType_Cos:
		pairs = append(pairs, cos.WithStorageFeatures(cos.StorageFeatures{
			VirtualDir: true,
		}))
		return cos.NewStorager(pairs...)
	case StorageType_Dropbox:
		return dropbox.NewStorager(pairs...)
	case StorageType_Fs:
		return fs.NewStorager(pairs...)
	case StorageType_Gcs:
		pairs = append(pairs, gcs.WithStorageFeatures(gcs.StorageFeatures{
			VirtualDir: true,
		}))
		return gcs.NewStorager(pairs...)
	case StorageType_Kodo:
		pairs = append(pairs, kodo.WithStorageFeatures(kodo.StorageFeatures{
			VirtualDir: true,
		}))
		return kodo.NewStorager(pairs...)
	case StorageType_Oss:
		pairs = append(pairs, oss.WithStorageFeatures(oss.StorageFeatures{
			VirtualDir: true,
		}))
		return oss.NewStorager(pairs...)
	case StorageType_Qingstor:
		pairs = append(pairs, qingstor.WithStorageFeatures(qingstor.StorageFeatures{
			VirtualDir: true,
		}))
		return qingstor.NewStorager(pairs...)
	case StorageType_S3:
		// always enable path style for s3, to avoid dns problem
		pairs = append(pairs, s3.WithForcePathStyle(true))
		return s3.NewStorager(pairs...)
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
