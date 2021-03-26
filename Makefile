SHELL := /bin/bash
GO_BUILD_OPTION := -trimpath -tags netgo

.PHONY: check format vet lint build test

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check      to format, vet and lint "
	@echo "  build      to create bin directory and build dm"

check: format vet

format:
	@echo "go fmt"
	@go fmt ./...
	@echo "ok"

generate:
	@echo "generate code"
	@go generate ./...
	@echo "ok"

vet:
	@echo "go vet"
	@go vet ./...
	@echo "ok"

build: tidy check
	@echo "build dm"
	@go build ${GO_BUILD_OPTION} -race -o ./bin/dm
	@echo "ok"

test:
	@echo "run test"
	@go test -gcflags=-l -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	@go tool cover -html="coverage.txt" -o "coverage.html"
	@echo "ok"

tidy:
	@echo "Tidy and check the go mod files"
	@go mod tidy
	@go mod verify
	@echo "Done"