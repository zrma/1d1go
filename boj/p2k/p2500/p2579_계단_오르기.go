package p2500

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve2579(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	if n == 1 {
		_, _ = fmt.Fprint(writer, arr[0])
		return
	}
	if n == 2 {
		_, _ = fmt.Fprint(writer, arr[0]+arr[1])
		return
	}

	dp := make([]int, n)

	dp[0] = arr[0]
	dp[1] = arr[0] + arr[1]
	dp[2] = integer.Max(arr[0], arr[1]) + arr[2]

	for i := 3; i < n; i++ {
		dp[i] = integer.Max(dp[i-2], dp[i-3]+arr[i-1]) + arr[i]
	}
	_, _ = fmt.Fprint(writer, dp[n-1])
}
