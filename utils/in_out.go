package utils

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
