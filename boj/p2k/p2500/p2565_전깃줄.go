package p2500

import (
	"fmt"
	"sort"

	"1d1go/utils/integer"
)

func Solve2565(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i][0], &arr[i][1])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	dp := make([]int, n)
	for i := range arr {
		dp[i] = arr[i][1]
	}

	longestIncSubSeqLen := integer.LongestIncSubSeqLenSquare(dp)
	res := n - longestIncSubSeqLen
	_, _ = fmt.Fprint(writer, res)
}
