FROM golang:1.22-bullseye

RUN groupadd nonroot \
    && useradd -ms /bin/bash nonroot -g nonroot

USER nonroot

COPY . /go/src/1d1go
WORKDIR /go/src/1d1go

RUN go mod download

RUN go fmt ./...
RUN go vet ./...
RUN go test -race -covermode=atomic $(go list ./... | grep -v /cmd/)
