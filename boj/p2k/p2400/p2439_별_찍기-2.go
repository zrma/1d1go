package p2400

import (
	"fmt"
	"io"
)

func Solve2439(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-1-i {
				_, _ = fmt.Fprint(writer, " ")
			} else {
				_, _ = fmt.Fprint(writer, "*")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
