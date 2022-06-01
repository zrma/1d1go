package p25000

type Writer interface {
	Write(p []byte) (n int, err error)
	Flush() error
}
