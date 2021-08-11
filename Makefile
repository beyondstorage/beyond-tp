SHELL := /bin/bash
GO_BUILD_OPTION := -trimpath -tags netgo

.PHONY: check format vet lint build test

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check      to format, vet and lint "
	@echo "  build      to create bin directory and build beyondtp"

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

build-frontend:
	@echo "build frontend"
	@git clean -Xf api/ui/
	cd ./ui && flutter build web --release
	@cp -r ui/build/web/* api/ui
	@echo "ok"

build: tidy check build-frontend
	@echo "build beyondtp"
	go build ${GO_BUILD_OPTION} -race -o ./bin/beyondtp ./cmd/beyondtp
	@echo "ok"

release: generate tidy check build-frontend
	@echo "release beyondtp"
	@-rm ./releases/*
	@mkdir -p ./releases

	@echo "build for linux amd64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/linux-amd64/beyondtp ./cmd/beyondtp
	tar -C ./bin/linux-amd64/ -czf ./releases/beyondtp_linux_amd64.tar.gz beyondtp

	@echo "build for macos amd64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/darwin-amd64/beyondtp ./cmd/beyondtp
	tar -C ./bin/darwin-amd64/ -czf ./releases/beyondtp_darwin_amd64.tar.gz beyondtp

	@echo "build for macos arm64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build ${GO_BUILD_OPTION} -o ./bin/darwin-arm64/beyondtp ./cmd/beyondtp
	tar -C ./bin/darwin-arm64/ -czf ./releases/beyondtp_darwin_arm64.tar.gz beyondtp

	@echo "build for windows amd64"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/windows-amd64/beyondtp ./cmd/beyondtp
	tar -C ./bin/windows-amd64/ -czf ./releases/beyondtp_windows_amd64.tar.gz beyondtp

test:
	@echo "run test"
	@go test -race -v -count=1 ./...
	@echo "ok"

tidy:
	@echo "Tidy and check the go mod files"
	@go mod tidy
	@go mod verify
	@echo "Done"