package p14400

import (
	"fmt"
)

func Solve14425(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	exist := make(map[string]bool)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		exist[s] = true
	}

	cnt := 0
	for i := 0; i < m; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		if exist[s] {
			cnt++
		}
	}
	_, _ = fmt.Fprint(writer, cnt)
}
