package p9300

import (
	"fmt"
	"io"
)

func Solve9372(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, m int
		_, _ = fmt.Fscan(reader, &n, &m)

		for j := 0; j < m; j++ {
			var a, b int
			_, _ = fmt.Fscan(reader, &a, &b)
		}

		_, _ = fmt.Fprintln(writer, n-1)
	}
}
