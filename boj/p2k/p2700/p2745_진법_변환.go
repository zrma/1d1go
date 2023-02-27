package p2700

import (
	"fmt"
	"io"
)

func Solve2745(reader io.Reader, writer io.Writer) {
	var s string
	var b int
	_, _ = fmt.Fscanln(reader, &s, &b)

	var n int
	for _, c := range s {
		n *= b
		if c >= '0' && c <= '9' {
			n += int(c - '0')
		} else {
			n += int(c - 'A' + 10)
		}
	}
	_, _ = fmt.Fprint(writer, n)
}
