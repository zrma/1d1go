package p2500

//go:generate go run github.com/golang/mock/mockgen -destination=./mocks/mock_in_out.go -package=mocks . InOut

type InOut interface {
	Scan(a ...any) (n int, err error)
	Println(a ...any) (n int, err error)
}

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
