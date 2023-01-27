package p2600

import (
	"fmt"
	"io"
)

func Solve2675(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var r int
		_, _ = fmt.Fscan(reader, &r)

		var s string
		_, _ = fmt.Fscan(reader, &s)
		for _, c := range s {
			for j := 0; j < r; j++ {
				_, _ = fmt.Fprint(writer, string(c))
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
