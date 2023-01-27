package p17400

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve17404(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	costs := make([][3]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &costs[i][0], &costs[i][1], &costs[i][2])
	}

	const inf = 1_234_567_891
	min := inf
	for i := 0; i < 3; i++ {
		dp := make([][3]int, n)
		for j := 0; j < 3; j++ {
			if i == j {
				dp[0][j] = costs[0][j]
			} else {
				dp[0][j] = inf
			}
		}

		for j := 1; j < n; j++ {
			dp[j][0] = integer.Min(dp[j-1][1], dp[j-1][2]) + costs[j][0]
			dp[j][1] = integer.Min(dp[j-1][0], dp[j-1][2]) + costs[j][1]
			dp[j][2] = integer.Min(dp[j-1][0], dp[j-1][1]) + costs[j][2]
		}

		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			min = integer.Min(min, dp[n-1][j])
		}
	}
	_, _ = fmt.Fprint(writer, min)
}
