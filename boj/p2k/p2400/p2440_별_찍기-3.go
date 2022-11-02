package p2400

import (
	"fmt"
)

func Solve2440(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			_, _ = fmt.Fprint(writer, "*")
		}
		_, _ = fmt.Fprintln(writer)
	}
}
