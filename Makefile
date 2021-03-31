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

build-frontend:
	@echo "build frontend"
	@git clean -Xf api/ui/
	cd ./ui && flutter build web --release
	@cp -r ui/build/web/* api/ui
	@echo "ok"

build: generate tidy check build-frontend
	@echo "build dm"
	go build ${GO_BUILD_OPTION} -race -o ./bin/dm ./cmd/dm
	@echo "ok"

release: generate tidy check build-frontend
	@echo "release dm"
	@-rm ./releases/*
	@mkdir -p ./releases

	@echo "build for linux amd64"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/linux-amd64/dm ./cmd/dm
	tar -C ./bin/linux-amd64/ -czf ./releases/dm_linux_amd64.tar.gz dm

	@echo "build for macos amd64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/darwin-amd64/dm ./cmd/dm
	tar -C ./bin/darwin-amd64/ -czf ./releases/dm_darwin_amd64.tar.gz dm

	@echo "build for macos arm64"
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build ${GO_BUILD_OPTION} -o ./bin/darwin-arm64/dm ./cmd/dm
	tar -C ./bin/darwin-arm64/ -czf ./releases/dm_darwin_arm64.tar.gz dm

	@echo "build for windows amd64"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${GO_BUILD_OPTION} -o ./bin/windows-amd64/dm ./cmd/dm
	tar -C ./bin/windows-amd64/ -czf ./releases/dm_windows_amd64.tar.gz dm

test:
	@echo "run test"
	@go test -race -v -count=1 ./...
	@echo "ok"

tidy:
	@echo "Tidy and check the go mod files"
	@go mod tidy
	@go mod verify
	@echo "Done"