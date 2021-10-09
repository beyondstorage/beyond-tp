package task

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

const (
	LargeFileBoundary = 1024 * 1024 * 1024 // Use 1 GiB as large file boundary
)

func MustMarshal(in interface{}) []byte {
	bs, err := msgpack.Marshal(in)
	if err != nil {
		panic(fmt.Errorf("marshal: %w", err))
	}
	return bs
}

func MustUnmarshal(bs []byte, in interface{}) {
	err := msgpack.Unmarshal(bs, in)
	if err != nil {
		panic(fmt.Errorf("unmarshal: %w", err))
	}
}
