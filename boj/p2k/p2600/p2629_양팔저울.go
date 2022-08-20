package p2600

import (
	"fmt"
)

func Solve2629(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const (
		offset      = 40_000
		maxLen      = 40_000
		maxWeight   = 500
		maxCount    = 30
		weightLimit = maxWeight * maxCount
	)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, offset+maxLen+1)
	}
	dp[0][offset] = true

	for i := 1; i <= n; i++ {
		var weight int
		_, _ = fmt.Fscan(reader, &weight)

		for j := 0; j <= offset+weightLimit+1; j++ {
			if !dp[i-1][j] {
				continue
			}
			dp[i][j] = true
			if j+weight <= offset+weightLimit {
				dp[i][j+weight] = true
			}
			if j-weight >= 0 {
				dp[i][j-weight] = true
			}
		}
	}

	var m int
	_, _ = fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var weight int
		_, _ = fmt.Fscan(reader, &weight)
		if dp[n][weight+offset] {
			_, _ = fmt.Fprint(writer, "Y")
		} else {
			_, _ = fmt.Fprint(writer, "N")
		}
		if i < m-1 {
			_, _ = fmt.Fprint(writer, " ")
		}
	}
}
