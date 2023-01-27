package p2200

import (
	"fmt"
	"io"
)

func Solve2293(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	dp := make([]int, k+1)
	dp[0] = 1

	for _, v := range arr {
		for j := 1; j < k+1; j++ {
			if j-v >= 0 {
				dp[j] += dp[j-v]
			}
		}
	}
	_, _ = fmt.Fprint(writer, dp[k])
}
