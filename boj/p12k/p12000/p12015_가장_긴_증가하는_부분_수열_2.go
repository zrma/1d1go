package p12000

import (
	"fmt"
	"sort"
)

func Solve12015(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	dp := make([]int, 1, n)
	_, _ = fmt.Fscan(reader, &dp[0])

	for i := 1; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if dp[len(dp)-1] < v {
			dp = append(dp, v)
			continue
		}

		idx := sort.SearchInts(dp, v)
		dp[idx] = v
	}

	_, _ = fmt.Fprint(writer, len(dp))
}
