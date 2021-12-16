//go:build tools

package main

import (
	_ "go-micro.dev/v4/cmd/protoc-gen-micro@v4"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1"
)
