package p1400

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve1463(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	dp := MoveToOne(n)
	
	_, _ = fmt.Fprint(writer, dp[n])
}

func MoveToOne(n int) []int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = integer.Min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = integer.Min(dp[i], dp[i/3]+1)
		}
	}
	return dp
}
