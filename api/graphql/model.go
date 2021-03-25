package graphql

import (
	"github.com/aos-dev/noah/proto"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

func (t *CreateTask) FormatTask() (*proto.Task, error) {
	// opt are task-level options, always handled as map
	opt := make(map[string]interface{})
	if t.Options != nil {
		opt = t.Options.(map[string]interface{})
	}

	// TODO: conduct other tasks, such as move or sync
	copyFileJob := &proto.CopyDir{
		Src:       0,
		Dst:       1,
		SrcPath:   "",
		DstPath:   "",
		Recursive: opt["recursive"].(bool),
	}
	content, err := protobuf.Marshal(copyFileJob)
	if err != nil {
		return nil, err
	}

	copyFileTask := &proto.Task{
		Id: uuid.NewString(),
		Endpoints: []*proto.Endpoint{
			t.Src.parse(),
			t.Dst.parse(),
		},
		Job: &proto.Job{
			Id:      uuid.NewString(),
			Type:    uint32(t.Type),
			Content: content,
		},
	}
	return copyFileTask, nil
}

// parse Endpoint into proto.Endpoint
func (e *Endpoint) parse() *proto.Endpoint {
	// ensure options handled as map[string]interface{}
	opt := make(map[string]interface{})
	if e.Options != nil {
		opt = e.Options.(map[string]interface{})
	}
	pairs := make([]*proto.Pair, 0, len(opt)+1) // +1 for work dir inject

	// conduct pairs with endpoint's options
	for k, v := range opt {
		pairs = append(pairs, &proto.Pair{Key: k, Value: v.(string)})
	}

	// inject work dir into pairs with given path
	pairs = append(pairs, &proto.Pair{Key: "work_dir", Value: e.Path})

	return &proto.Endpoint{Type: e.Type.String(), Pairs: pairs}
}
