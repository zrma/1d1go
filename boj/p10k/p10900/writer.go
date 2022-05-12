package p10900

type Writer interface {
	Write(p []byte) (n int, err error)
}

//go:generate go run github.com/golang/mock/mockgen -destination=./writer_mock.go -package=p10900 . Writer
