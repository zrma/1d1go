# Helloworld Service

This is the Helloworld service

Generated with

```
micro new helloworld
```

## Requirement

```
# protoc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

(linux) brew install protobuf
(windows) scoop install protobuf

# protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

# protoc-gen-micro
go install go-micro.dev/v4/cmd/protoc-gen-micro@v4
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- Alias: helloworld

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service

```
./helloworld
```

Build a docker image

```
make docker
```
