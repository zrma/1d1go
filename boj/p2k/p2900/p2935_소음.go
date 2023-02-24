package p2900

import (
	"fmt"
	"io"
)

func Solve2935(reader io.Reader, writer io.Writer) {
	var a, op, b string
	_, _ = fmt.Fscan(reader, &a, &op, &b)

	n := len(a)
	m := len(b)

	if op == "*" {
		_, _ = fmt.Fprint(writer, "1")
		for i := 0; i < n+m-2; i++ {
			_, _ = fmt.Fprint(writer, "0")
		}
		return
	}

	if n == m {
		_, _ = fmt.Fprint(writer, "2")
		for i := 0; i < n-1; i++ {
			_, _ = fmt.Fprint(writer, "0")
		}
		return
	}

	if n > m {
		_, _ = fmt.Fprint(writer, a[:n-m])
		_, _ = fmt.Fprint(writer, b)
	} else if n < m {
		_, _ = fmt.Fprint(writer, b[:m-n])
		_, _ = fmt.Fprint(writer, a)
	}
}
