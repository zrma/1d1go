package p5200

import (
	"fmt"
	"io"
)

func Solve5218(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var a, b string
		_, _ = fmt.Fscanln(reader, &a, &b)

		_, _ = fmt.Fprintf(writer, "Distances: ")
		for j := 0; j < len(a); j++ {
			if a[j] <= b[j] {
				_, _ = fmt.Fprint(writer, b[j]-a[j])
			} else {
				_, _ = fmt.Fprint(writer, b[j]-a[j]+26)
			}

			if j != len(a)-1 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
