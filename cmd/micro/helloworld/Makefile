
GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc --go_out=proto --go_opt=paths=source_relative \
        --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
        --micro_out=proto --micro_opt=paths=source_relative \
        -I proto helloworld.proto

.PHONY: build
build:
	go build -o helloworld *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t helloworld:latest
