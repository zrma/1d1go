package p2500

import (
	"fmt"
	"sort"
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

	dp := make([]int, 1, n)
	dp[0] = arr[0][1]

	for i := 1; i < n; i++ {
		v := arr[i][1]

		if dp[len(dp)-1] < v {
			dp = append(dp, v)
			continue
		}

		idx := sort.SearchInts(dp, v)
		dp[idx] = v
	}

	res := n - len(dp)
	_, _ = fmt.Fprint(writer, res)
}
