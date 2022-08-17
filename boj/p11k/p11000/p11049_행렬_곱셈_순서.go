package p11000

import (
	"fmt"
)

func Solve11049(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int64, n+1)
	_, _ = fmt.Fscan(reader, &arr[0], &arr[1])
	for i := 1; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i], &arr[i+1])
	}

	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	const maxVal int64 = 999_999_999_999

	for end := 1; end < n; end++ {
		for begin := 1; begin+end <= n; begin++ {
			dp[begin][begin+end] = maxVal

			for k := begin; k < begin+end; k++ {
				v := dp[begin][k] + dp[k+1][begin+end] + arr[begin-1]*arr[k]*arr[begin+end]
				if v < dp[begin][begin+end] {
					dp[begin][begin+end] = v
				}
			}
		}
	}
	_, _ = fmt.Fprint(writer, dp[1][n])
}
