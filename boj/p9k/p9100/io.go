package p9100

type Reader interface {
	ReadString(delim byte) (string, error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
