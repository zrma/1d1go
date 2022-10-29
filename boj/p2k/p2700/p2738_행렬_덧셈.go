package p2700

import (
	"fmt"
)

func Solve2738(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Fscan(reader, &mat[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var v int
			_, _ = fmt.Fscan(reader, &v)
			mat[i][j] += v
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				_, _ = fmt.Fprintln(writer, mat[i][j])
			} else {
				_, _ = fmt.Fprint(writer, mat[i][j], " ")
			}
		}
	}
}
