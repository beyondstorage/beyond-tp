package models

//go:generate protoc --go_out=. --go_opt=paths=source_relative job.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative task.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative identity.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative agent.proto
