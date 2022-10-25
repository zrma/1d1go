package p14000

import (
	"fmt"

	"1d1go/boj/p11k/p11000"
)

func Solve14002(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
		dp[i] = 1
	}

	max := p11000.LongestIncreasingSubsequence(n, arr, dp)
	_, _ = fmt.Fprintln(writer, max)

	res := make([]int, max)
	for i := n - 1; i >= 0; i-- {
		if dp[i] == max {
			res[max-1] = arr[i]
			max--
		}
	}

	for i := 0; i < len(res); i++ {
		_, _ = fmt.Fprintf(writer, "%d ", res[i])
	}
}
