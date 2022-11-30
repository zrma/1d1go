package p11700

import (
	"fmt"
)

func Solve11727(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const mod = 10_007
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + 2*dp[i-2]) % mod
	}

	_, _ = fmt.Fprint(writer, dp[n])
}
