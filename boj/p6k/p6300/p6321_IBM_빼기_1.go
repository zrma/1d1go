package p6300

import (
	"fmt"
	"io"
)

func Solve6321(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	for i := 1; i <= n; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)

		_, _ = fmt.Fprintf(writer, "String #%d\n", i)
		for _, c := range s {
			if c == 'Z' {
				_, _ = fmt.Fprintf(writer, "A")
			} else {
				_, _ = fmt.Fprintf(writer, "%c", c+1)
			}
		}
		_, _ = fmt.Fprintln(writer)
		_, _ = fmt.Fprintln(writer)
	}
}
