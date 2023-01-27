FROM golang:1.19.5-bullseye

COPY . /go/src/1d1go
WORKDIR /go/src/1d1go

RUN go mod download

RUN go fmt ./...
RUN go vet ./...
RUN go test -race -covermode=atomic $(go list ./... | grep -v /cmd/)
