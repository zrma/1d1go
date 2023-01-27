package p11700

import (
	"fmt"
	"io"
)

func Solve11726(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const max = 1000 + 4
	dp := make([]int, max)

	dp[1] = 1
	dp[2] = 2
	for i := 3; i < max; i++ {
		dp[i] = (dp[i-2] + dp[i-1]) % 10_007
	}
	_, _ = fmt.Fprint(writer, dp[n])
}
