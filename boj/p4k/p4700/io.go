package p4700

type Reader interface {
	Read(p []byte) (n int, err error)
	ReadString(delim byte) (string, error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
