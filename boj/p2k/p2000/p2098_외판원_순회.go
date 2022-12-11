package p2000

import (
	"fmt"
)

func Solve2098(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	costs := make([][]int, n)
	for i := 0; i < n; i++ {
		costs[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &costs[i][j])
		}
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = -1
		}
	}

	res := solve2098WithDFS(0, 1, n, costs, dp)
	_, _ = fmt.Fprint(writer, res)
}

func solve2098WithDFS(cur, visited, n int, costs, dp [][]int) int {
	const inf = 1_234_567_891

	if visited == (1<<n)-1 {
		if costs[cur][0] > 0 {
			return costs[cur][0]
		}
		return inf
	}

	if dp[cur][visited] != -1 {
		return dp[cur][visited]
	}

	res := inf
	for i := 0; i < n; i++ {
		if visited&(1<<i) > 0 {
			continue
		}

		if costs[cur][i] == 0 {
			continue
		}

		if cost := solve2098WithDFS(i, visited|(1<<i), n, costs, dp) + costs[cur][i]; cost < res {
			res = cost
		}
	}

	dp[cur][visited] = res
	return res
}
