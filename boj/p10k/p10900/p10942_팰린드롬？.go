package p10900

import (
	"fmt"
	"io"
)

func Solve10942(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n+1)
	dp := make([][]int, n+1)
	dp[0] = make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
		dp[i] = make([]int, n+1)
		dp[i][i] = 1

		if i > 1 && arr[i] == arr[i-1] {
			dp[i-1][i] = 1
		}
	}

	for length := 2; length < n; length++ {
		for start := 1; start+length <= n; start++ {
			end := start + length
			if arr[start] == arr[end] && dp[start+1][end-1] > 0 {
				dp[start][end] = 1
			}
		}
	}

	var m int
	_, _ = fmt.Fscan(reader, &m)
	for i := 0; i < m; i++ {
		var start, end int
		_, _ = fmt.Fscan(reader, &start, &end)
		_, _ = fmt.Fprintln(writer, dp[start][end])
	}
}
