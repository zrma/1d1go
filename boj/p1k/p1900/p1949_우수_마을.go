package p1900

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve1949(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	people := make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Fscan(reader, &people[i])
	}

	adj := make([][]int, n+1)
	for i := 1; i <= n-1; i++ {
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)

		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	dp := make([][2]int, n+1)
	getDp(adj, dp, people, 1, 0)

	_, _ = fmt.Fprint(writer, integer.Max(dp[1][0], dp[1][1]))
}

func getDp(adj [][]int, dp [][2]int, people []int, cur, parent int) {
	dp[cur][0] = 0 // not contain
	dp[cur][1] = people[cur]

	for _, next := range adj[cur] {
		if next == parent {
			continue
		}

		getDp(adj, dp, people, next, cur)

		dp[cur][0] += integer.Max(dp[next][0], dp[next][1]) // cur 빼면 next (포함/빼는거) 중 큰 것
		dp[cur][1] += dp[next][0]                           // cur 넣으면 next 무조건 뺌
	}
}
