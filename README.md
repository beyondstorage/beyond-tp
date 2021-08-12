# beyond-tp

[![go version](https://img.shields.io/github/go-mod/go-version/beyondstorage/beyond-tp)]()
[![license](https://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/beyondstorage/beyond-tp/blob/master/LICENSE)
[![unit test status](https://github.com/beyondstorage/beyond-tp/workflows/Unit%20Test/badge.svg?branch=master)](https://github.com/beyondstorage/beyond-tp/actions?query=workflow%3A%22Unit+Test%22)
[![beyond-tp dev](https://img.shields.io/matrix/beyondstorage@beyond-tp:matrix.org.svg?label=beyond-tp&logo=matrix)](https://matrix.to/#/#beyondstorage@beyond-tp:matrix.org)

beyond-tp is a data migration service.

**This project is now under high development, please do not apply to production environment**

- [Contributor Guide](#contributor-guide)
- [Development Requirement](#development-requirement)
  - [protobuf](#protobuf)
  - [flutter](#flutter)
- [Build](#build)
  - [Build Binary](#build-binary)
  - [Build Frontend Pages](#build-frontend-pages)
  - [Generate Code](#generate-code)
- [Usage](#usage)  
  - [Start a Server](#start-a-server)
  - [Start a Staff](#start-a-staff)
  - [What's more](#whats-more)

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

## Build

### Build Binary

```
make build
```

Run `make build` to build the binary file used to start a beyond-tp server or staff.

### Build Frontend Pages

```
make build-frontend
```

Run `make build-frontend` to build frontend pages. 

**Notice** that beyond-tp uses flutter to build web pages,
so please make sure flutter is well [installed](#flutter), and added in `$PATH` directories.

For more details, please run `make help` as a reference.

### Generate Code

```
make generate
```

Run `make generate` to generate the latest code by GraphQL schema and grpc `.proto` files.

**Notice** that beyond-tp uses protobuf tools to generate grpc code,
so please make sure protobuf is well [installed](#protobuf), 
and added in `$PATH` directories for the protocol buffer compiler to find it.

## Usage

### Start a Server

`Server` used as a web server as well as a task manager. It allows you to manage beyond-tp visually,
and support a rpc server to communicate with `Staffs`. You can easily start a beyond-tp server by:

```
beyondtp server --db /path/to/db --host localhost --port 7436 --rpc-port 7000
```

This command will start a beyond-tp server, including a web server, with local db indicated by `db` flag. 
You can visit the home page at `localhost:7436/ui` in your web browser.

Run `beyondtp server --help` for more examples and flags' usage.

### Start a Staff

`Staff` used as a task executor. It can run a migration task, distributed by `Server`. 
You can easily start a beyond-tp staff by:

```
beyondtp staff --host localhost --manager localhost:7000
```

This command will start a beyond-tp staff connected to the manager `localhost:7000`.
The manager's host and port are specified in `server` command by flag `host` and `rpc-port`.

Run `beyondtp staff --help` for more examples and flags' usage.

### What's more

We support to use environment variable to replace flags. All the flags follow these rules:

- All flags correspond to environment variable with prefix `BEYONDTP_`
- Kebab-case (`log-level`) to Snake_case (`LOG_LEVEL`) with upper case
- Flags for specific commands should add command prefix, e.g. `BEYONDTP_SERVER_PORT` and `BEYONDTP_STAFF_PORT` for 
  `--port` flag under `server` command and `staff` command 
- Global flags do not need to add command prefix, e.g. `BEYONDTP_LOG_LEVEL`

If the same flag set by both environment variable and command flag, command flag will be 
higher priority than environment variable, for example:

```
BEYONDTP_LOG_LEVEL=debug bin/beyondtp server --db badger --log-level=warn
```

The final log level would be `warn` instead of `debug`.

For more command usage, please run `beyondtp --help` as a reference.

## Useful links

- [Project Roadmap](https://github.com/orgs/beyondstorage/projects/2#card-60949498)
- [Discussions on forum](https://forum.beyondstorage.io/c/development/dm/10)
