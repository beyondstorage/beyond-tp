# beyond-tp

[![beyond-tp dev](https://img.shields.io/matrix/beyondstorage@beyond-tp:matrix.org.svg?label=beyond-tp&logo=matrix)](https://matrix.to/#/#beyondstorage@beyond-tp:matrix.org)

beyond-tp is a data migration service.

**This project is now under high development, please do not apply to production environment**

## Contributor Guide

You are always welcome to act as a contributor to beyond-tp. 

Since beyond-tp is an application under [BeyondStorage](https://beyondstorage.io), 
you should follow the contributor guide of the community.

Please feel free to submit issues, create a pull request, or just star this repo :)

For more details, please refer to: <https://beyondstorage.io/docs/general/contributor-guide>

## Development Requirement

### protobuf

Beyond-tp uses protobuf for RPC call and data serialization. So protobuf tools are needed when generate code
from `.proto` files, related components are listed below:

#### protoc

`protoc` is the `Protocol Compiler`. To install `protoc`, you can
take [Protocol Compiler Installation](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation)
as a reference.

#### protoc-gen-go & protoc-gen-go-grpc

`protoc-gen-go` and `protoc-gen-go-grpc` are two plugins used to generate Go code from `.proto` files.

Install them using Go 1.16 or higher by running:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

This will install `protoc-gen-go` and `protoc-gen-go-grpc` binary files in `$GOBIN`. 
Set the `$GOBIN` environment variable to change the installation location. 
It must be in your `$PATH` for the protocol buffer compiler to find it.

You can also specify a version number by replacing `latest` after `@`.

Since we develop beyond-tp with:

```
➜  ~ protoc --version
libprotoc 3.15.7
➜  ~ protoc-gen-go --version
protoc-gen-go v1.25.0
➜  ~ protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.1.0
```

So we recommend you to use these tools equal or higher than this.

### flutter

Beyond-tp uses flutter to conduct frontend web pages. `v2.0.1` or higher version is recommended for development.

To install flutter, you can take [flutter installation](https://flutter.dev/docs/get-started/install) as a reference.

**Notice**: For some network reason, you can also get help from <https://flutter.cn/community/china>.

## Useful links

- [Project Roadmap](https://github.com/orgs/beyondstorage/projects/2#card-60949498)
- [Discussions on forum](https://forum.beyondstorage.io/c/development/dm/10)
