package p14000

import (
	"fmt"
)

func Solve14002(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
		dp[i] = 1
	}

	max := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] <= arr[j] {
				continue
			}
			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				if max < dp[i] {
					max = dp[i]
				}
			}
		}
	}

	_, _ = fmt.Fprintln(writer, max)

	res := make([]int, max)
	for i := n - 1; i >= 0; i-- {
		if dp[i] == max {
			res[max-1] = arr[i]
			max--
		}
	}

	for i := 0; i < len(res); i++ {
		_, _ = fmt.Fprintf(writer, "%d ", res[i])
	}
}
