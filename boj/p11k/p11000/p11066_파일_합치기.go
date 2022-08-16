package p11000

import (
	"fmt"
)

func Solve11066(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		mergeFile(reader, writer)
	}
}

func mergeFile(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int64, n+1)
	sum := make([]int64, n+1)
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
		sum[i] = sum[i-1] + arr[i]
	}

	const maxVal int64 = 999_999_999_999

	for end := 1; end <= n; end++ {
		for begin := 1; begin+end <= n; begin++ {
			dp[begin][begin+end] = maxVal

			for k := begin; k < begin+end; k++ {
				v := dp[begin][k] + dp[k+1][begin+end] + sum[begin+end] - sum[begin-1]
				if v < dp[begin][begin+end] {
					dp[begin][begin+end] = v
				}
			}
		}
	}
	_, _ = fmt.Fprintln(writer, dp[1][n])
}
