package p2500

import (
	"fmt"

	"1d1go/utils/integer"
)

func Solve2533(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	adj := make([][]int, n+1)
	for i := 1; i < n; i++ {
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	dp := make([][2]int, n+1)
	getDp(adj, dp, 1, 0)

	_, _ = fmt.Fprint(writer, integer.Min(dp[1][0], dp[1][1]))
}

func getDp(adj [][]int, dp [][2]int, cur, parent int) {
	dp[cur][0] = 0 // not contain
	dp[cur][1] = 1 // contain

	for _, next := range adj[cur] {
		if next == parent {
			continue
		}

		getDp(adj, dp, next, cur)

		dp[cur][0] += dp[next][1]                           // cur 빼면 next  무조건 포함
		dp[cur][1] += integer.Min(dp[next][0], dp[next][1]) // cur 넣으면 next (포함/빼는거) 중 작은 것
	}
}
