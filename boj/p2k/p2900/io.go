package p2900

type Reader interface {
	Read(p []byte) (n int, err error)
	ReadString(delim byte) (string, error)
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
