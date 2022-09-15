package p6000

import (
	"fmt"
)

func Solve6064(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var m, n, x, y int
		_, _ = fmt.Fscan(reader, &m, &n, &x, &y)
		_, _ = fmt.Fprintln(writer, solve6064(m, n, x, y))
	}
}

func solve6064(m, n, x, y int) int {
	x -= 1
	y -= 1
	for i := x; i <= m*n; i += m {
		if i%n == y {
			return i + 1
		}
	}
	return -1
}
