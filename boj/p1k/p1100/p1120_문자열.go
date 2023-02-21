package p1100

import (
	"fmt"
	"io"
)

func Solve1120(reader io.Reader, writer io.Writer) {
	var a, b string
	_, _ = fmt.Fscan(reader, &a, &b)

	n := len(a)
	m := len(b)

	min := 50
	for i := 0; i <= m-n; i++ {
		diff := 0
		for j := 0; j < n; j++ {
			if a[j] != b[i+j] {
				diff++
			}
		}
		if diff < min {
			min = diff
		}
	}

	_, _ = fmt.Fprint(writer, min)
}
