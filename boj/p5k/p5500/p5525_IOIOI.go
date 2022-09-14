package p5500

import (
	"fmt"
)

func Solve5525(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n)
	_, _ = fmt.Fscan(reader, &m)

	var s string
	_, _ = fmt.Fscan(reader, &s)

	res := 0
	dp := make([]int, m)
	for i := 2; i < m; i++ {
		if s[i] == 'I' && s[i-1] == 'O' && s[i-2] == 'I' {
			dp[i] = dp[i-2] + 1
			if dp[i] >= n {
				res++
			}
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
