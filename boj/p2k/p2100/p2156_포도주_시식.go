package p2100

import (
	"fmt"
	"strconv"

	"1d1go/utils/integer"
)

func Solve2156(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	dp := make([]int, n)
	dp[0] = arr[0]
	if n == 1 {
		_, _ = fmt.Fprint(writer, dp[0])
		return
	}

	dp[1] = arr[0] + arr[1]
	if n == 2 {
		_, _ = fmt.Fprint(writer, dp[1])
		return
	}

	dp[2] = integer.Max(arr[0]+arr[1], arr[1]+arr[2], arr[0]+arr[2])
	if n == 3 {
		_, _ = fmt.Fprint(writer, dp[2])
		return
	}

	for i := 3; i < n; i++ {
		dp[i] = integer.Max(dp[i-3]+arr[i-1]+arr[i], dp[i-2]+arr[i], dp[i-1])
	}
	_, _ = fmt.Fprint(writer, dp[n-1])
}
