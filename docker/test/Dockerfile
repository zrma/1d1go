FROM golang:1.17.7-bullseye

COPY . /go/src/1d1go
WORKDIR /go/src/1d1go

RUN go mod download

RUN go test -race -covermode=atomic $(go list ./... | grep -v /cmd/)
