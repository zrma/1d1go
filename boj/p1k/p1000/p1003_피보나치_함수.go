package p1000

import (
	"fmt"
)

func Solve1003(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	const maxLen = 40 + 1
	var dp [maxLen][2]int
	dp[0] = [2]int{1, 0}
	dp[1] = [2]int{0, 1}
	for i := 2; i < maxLen; i++ {
		dp[i] = [2]int{dp[i-1][0] + dp[i-2][0], dp[i-1][1] + dp[i-2][1]}
	}

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)

		_, _ = fmt.Fprintln(writer, dp[n][0], dp[n][1])
	}
}
