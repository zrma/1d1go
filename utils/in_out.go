package utils

//go:generate go run github.com/golang/mock/mockgen -destination=./mocks/mock_writer.go -package=mocks . Writer

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
