package p11600

import (
	"fmt"
)

func Solve11660(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr2D := make([][]int, n+1)
	for i := range arr2D {
		arr2D[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			_, _ = fmt.Fscan(reader, &arr2D[i][j])
			arr2D[i][j] += arr2D[i-1][j] + arr2D[i][j-1] - arr2D[i-1][j-1]
		}
	}

	for i := 0; i < m; i++ {
		var x1, y1, x2, y2 int
		_, _ = fmt.Fscan(reader, &x1, &y1, &x2, &y2)

		res := arr2D[x2][y2] - arr2D[x2][y1-1] - arr2D[x1-1][y2] + arr2D[x1-1][y1-1]
		_, _ = fmt.Fprintln(writer, res)
	}
}
