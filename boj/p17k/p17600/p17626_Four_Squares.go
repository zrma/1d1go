package p17600

import (
	"fmt"
)

func Solve17626(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			if dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}

	_, _ = fmt.Fprint(writer, dp[n])
}
