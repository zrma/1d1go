package p2800

import (
	"fmt"
	"io"
)

func Solve2804(reader io.Reader, writer io.Writer) {
	var a, b string
	_, _ = fmt.Fscanln(reader, &a, &b)

	aIdx, bIdx := func() (int, int) {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					return i, j
				}
			}
		}
		return 0, 0
	}()

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(a); j++ {
			if i == bIdx {
				_, _ = fmt.Fprint(writer, string(a[j]))
			} else if j == aIdx {
				_, _ = fmt.Fprint(writer, string(b[i]))
			} else {
				_, _ = fmt.Fprint(writer, ".")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
