package p11900

import (
	"fmt"
	"io"
)

func Solve11945(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		for j := m - 1; j >= 0; j-- {
			_, _ = fmt.Fprint(writer, string(s[j]))
		}
		_, _ = fmt.Fprint(writer, "\n")
	}
}
