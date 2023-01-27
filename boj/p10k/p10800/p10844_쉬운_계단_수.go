package p10800

import (
	"fmt"
	"io"
)

func Solve10844(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const base = 10
	if n == 1 {
		_, _ = fmt.Fprint(writer, base-1) // exclude 0
		return
	}

	dp := make([][base]int64, n)
	for i := 1; i < base; i++ {
		dp[0][i] = 1
	}

	const mod = 1_000_000_000
	for i := 1; i < n; i++ {
		for j := 0; j < base; j++ {
			switch j {
			case 0:
				dp[i][j] = dp[i-1][1] % mod
			case base - 1:
				dp[i][j] = dp[i-1][base-2] % mod
			default:
				dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1]) % mod
			}
		}
	}

	var res int64
	for _, v := range dp[n-1] {
		res = (res + v) % mod
	}

	_, _ = fmt.Fprint(writer, res)
}
