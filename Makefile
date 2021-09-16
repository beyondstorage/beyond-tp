SHELL := /bin/bash
GO_BUILD_OPTION := -trimpath

.PHONY: check format vet lint build test

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check				to format code and vet check"
	@echo "  build				to create bin directory and build beyondtp"
	@echo "  build-frontend		to build flutter web and copy to api/ui directory"
	@echo "  generate			to generate code"
	@echo "  test				to run unit tests"
	@echo "  tidy				to tidy and check the go mod files"

check: format vet

format:
	go fmt ./...

generate:
	go generate ./...

vet:
	go vet ./...

build-frontend:
	cd ./ui && flutter build web --release
	cp -r ui/build/web/* api/ui

# We remove check target to work around build failed in container build.
# If we build in container, we will meet errors like:
#
# > go fmt
# > go: not formatting packages in dependency modules
# > package github.com/google/flatbuffers/grpc/tests: C++ source files not allowed when not using cgo or SWIG: grpctest.cpp message_builder_test.cpp
# > package github.com/google/flatbuffers/samples: C++ source files not allowed when not using cgo or SWIG: sample_bfbs.cpp sample_binary.cpp sample_text.cpp
# > package github.com/google/flatbuffers/tests: C++ source files not allowed when not using cgo or SWIG: monster_test.grpc.fb.cc native_type_test_impl.cpp test.cpp test_assert.cpp test_builder.cpp
# > make: *** [Makefile:19: format] Error 1
build: tidy
	CGO_ENABLED=1 go build ${GO_BUILD_OPTION} -race -o ./bin/beyondtp ./cmd/beyondtp

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
	go test -race -v -count=1 ./...

tidy:
	go mod tidy
	go mod verify
