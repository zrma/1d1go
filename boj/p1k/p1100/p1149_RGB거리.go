package p1100

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve1149(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const colorNum = 3
	arr := make([][colorNum]int, n)
	for i := range arr {
		for j := 0; j < colorNum; j++ {
			_, _ = fmt.Fscan(reader, &arr[i][j])
		}
	}

	dp := make([][colorNum]int, n)
	dp[0] = arr[0]

	for i := 1; i < n; i++ {
		for j := 0; j < colorNum; j++ {
			dp[i][j] = integer.Min(dp[i-1][(j+1)%colorNum], dp[i-1][(j+2)%colorNum]) + arr[i][j]
		}
	}
	_, _ = fmt.Fprint(writer, integer.Min(dp[n-1][0], dp[n-1][1], dp[n-1][2]))
}
