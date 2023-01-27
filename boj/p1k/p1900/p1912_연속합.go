package p1900

import (
	"fmt"
	"io"
)

func Solve1912(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	dp := make([]int, n)
	dp[0] = arr[0]
	maxVal := dp[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+arr[i], arr[i])
		maxVal = max(maxVal, dp[i])
	}

	_, _ = fmt.Fprint(writer, maxVal)
}
