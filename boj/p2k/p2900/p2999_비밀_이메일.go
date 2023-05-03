package p2900

import (
	"fmt"
	"io"
)

func Solve2999(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	r, c := getRC(len(s))

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			_, _ = fmt.Fprintf(writer, "%c", s[i+j*r])
		}
	}
}

func getRC(n int) (int, int) {
	maxR := 1
	for r := 1; r <= n; r++ {
		if n%r == 0 {
			c := n / r
			if r <= c {
				if maxR < r {
					maxR = r
				}
			}
		}
	}
	return maxR, n / maxR
}
