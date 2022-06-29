package p1200

import (
	"fmt"
)

func Solve1269(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		set[v] = true
	}

	unionCount := 0
	for i := 0; i < m; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		if _, ok := set[v]; ok {
			unionCount++
		}
	}
	_, _ = fmt.Fprint(writer, n+m-unionCount*2)
}
