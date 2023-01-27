package p25600

import (
	"fmt"
	"io"
)

func Solve25682(reader io.Reader, writer io.Writer) {
	var n, m, k int
	_, _ = fmt.Fscan(reader, &n, &m, &k)

	board0 := make([][]int, n+1)
	board1 := make([][]int, n+1)
	for i := range board0 {
		board0[i] = make([]int, m+1)
		board1[i] = make([]int, m+1)
	}

	const (
		w byte = 'W'
		b byte = 'B'
	)

	var (
		startWant, startOther = w, b
	)

	for i := 1; i <= n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		startWant, startOther = startOther, startWant
		want, other := startWant, startOther

		for j := 1; j <= m; j++ {
			want, other = other, want

			if s[j-1] == want {
				board0[i][j] = board0[i][j-1] + board0[i-1][j] - board0[i-1][j-1]
				board1[i][j] = board1[i][j-1] + board1[i-1][j] - board1[i-1][j-1] + 1
			} else {
				board0[i][j] = board0[i][j-1] + board0[i-1][j] - board0[i-1][j-1] + 1
				board1[i][j] = board1[i][j-1] + board1[i-1][j] - board1[i-1][j-1]
			}
		}
	}

	res := m * n
	for i := k; i <= n; i++ {
		for j := k; j <= m; j++ {
			diff0 := board0[i][j] - board0[i-k][j] - board0[i][j-k] + board0[i-k][j-k]
			diff1 := board1[i][j] - board1[i-k][j] - board1[i][j-k] + board1[i-k][j-k]

			if res > diff0 {
				res = diff0
			}
			if res > diff1 {
				res = diff1
			}
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
