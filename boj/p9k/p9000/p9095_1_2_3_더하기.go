package p9000

import (
	"fmt"
	"io"
)

func Solve9095(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	const max = 11 + 4
	dp := make([]int, max)

	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	// NOTE
	//   dp[4]
	//     case1: 3 + 1 -> dp[3] + 1
	//     case2: 2 + 2 -> dp[2] + 2
	//     case3: 1 + 3 -> dp[1] + 3
	//   dp[5]
	//     case1: 4 + 1 -> dp[4] + 1
	//     case2: 3 + 2 -> dp[3] + 2
	//     case3: 2 + 3 -> dp[2] + 3
	//   ...

	for i := 4; i < max-4; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)

		_, _ = fmt.Fprintln(writer, dp[n])
	}
}
