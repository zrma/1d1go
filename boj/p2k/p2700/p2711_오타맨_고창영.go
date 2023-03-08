package p2700

import (
	"fmt"
	"io"
)

func Solve2711(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		var s string
		_, _ = fmt.Fscanln(reader, &n, &s)
		_, _ = fmt.Fprintln(writer, s[:n-1]+s[n:])
	}
}
