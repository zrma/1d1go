package p10800

import (
	"fmt"
)

func Solve10816(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		counts[v]++
	}

	var m int
	_, _ = fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		_, _ = fmt.Fprint(writer, counts[v], " ")
	}
}
