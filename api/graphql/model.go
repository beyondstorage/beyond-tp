package graphql

import (
	"encoding/json"

	"github.com/aos-dev/noah/proto"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

type FsOption struct {
	Recursive bool `json:"recursive"`
}

func (t *CreateTask) FormatTask() (*proto.Task, error) {
	if err := t.Src.ParseOption(); err != nil {
		return nil, err
	}

	if err := t.Dst.ParseOption(); err != nil {
		return nil, err
	}

	copyFileJob := &proto.CopyDir{
		Src:       0,
		Dst:       1,
		SrcPath:   "",
		DstPath:   "",
		Recursive: t.Src.Option.(FsOption).Recursive,
	}
	content, err := protobuf.Marshal(copyFileJob)
	if err != nil {
		return nil, err
	}

	copyFileTask := &proto.Task{
		Id: uuid.NewString(),
		Endpoints: []*proto.Endpoint{
			{Type: t.Src.Type.String(), Pairs: []*proto.Pair{{Key: "work_dir", Value: t.Src.Path}}},
			{Type: t.Dst.Type.String(), Pairs: []*proto.Pair{{Key: "work_dir", Value: t.Dst.Path}}},
		},
		Job: &proto.Job{
			Id:      uuid.NewString(),
			Type:    uint32(t.Type),
			Content: content,
		},
	}
	return copyFileTask, nil
}

func (ep *Endpoint) ParseOption() error {
	switch ep.Type.String() {
	case "fs":
		res := FsOption{}
		data, err := json.Marshal(ep.Option)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(data, &res); err != nil {
			return err
		}
		ep.Option = res
	}
	return nil
}
