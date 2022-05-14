package utils

import (
	"fmt"
)

type InOut interface {
	Scan(a ...any) (n int, err error)
	Println(a ...any) (n int, err error)
}

func NewIO() InOut {
	return &impl{}
}

type impl struct{}

var (
	fnIn  = fmt.Scan
	fnOut = fmt.Println
)

func (i impl) Scan(args ...any) (n int, err error) {
	return fnIn(args...)
}

func (i impl) Println(args ...any) (n int, err error) {
	return fnOut(args...)
}

//go:generate go run github.com/golang/mock/mockgen -destination=./mocks/mock_scanner.go -package=mocks . Scanner
//go:generate go run github.com/golang/mock/mockgen -destination=./mocks/mock_writer.go -package=mocks . Writer

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
