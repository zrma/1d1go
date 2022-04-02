//go:build tools

package main

import (
	_ "go-micro.dev/v4/cmd/protoc-gen-micro"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
