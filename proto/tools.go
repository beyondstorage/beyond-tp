//+build tools

package proto

import (
	_ "github.com/golang/protobuf/protoc-gen-go/grpc"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
