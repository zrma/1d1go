package p10800

import (
	"fmt"
	"io"
)

func Solve10815(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	exist := make(map[int]bool)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		exist[v] = true
	}

	var m int
	_, _ = fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if exist[v] {
			_, _ = fmt.Fprint(writer, "1 ")
		} else {
			_, _ = fmt.Fprint(writer, "0 ")
		}
	}
}
