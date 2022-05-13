package utils

import (
	"fmt"
)

//go:generate go run github.com/golang/mock/mockgen -destination=./mock_in_out.go -package=utils . InOut

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
