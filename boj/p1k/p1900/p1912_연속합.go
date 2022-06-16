package p1900

import (
	"fmt"
	"strconv"
)

func Solve1912(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	dp := make([]int, n)
	dp[0] = arr[0]
	maxVal := dp[0]

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+arr[i], arr[i])
		maxVal = max(maxVal, dp[i])
	}

	_, _ = fmt.Fprint(writer, maxVal)
}
