package p8900

import (
	"fmt"
	"io"
)

func Solve8958(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		res := acc(s)
		_, _ = fmt.Fprintln(writer, res)
	}
}

func acc(s string) int {
	tot := 0
	cur := 1
	for _, c := range s {
		if c == 'O' {
			tot += cur
			cur++
		} else {
			cur = 1
		}
	}
	return tot
}
