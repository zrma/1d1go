package p2200

import (
	"fmt"
	"sort"

	"1d1go/utils/integer"
)

func Solve2213(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	weights := make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Fscan(reader, &weights[i])
	}

	adj := make([][]int, n+1)
	for i := 1; i <= n-1; i++ { // edge == node - 1
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)

		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	visited := make([]bool, n+1)
	visited[1] = true

	dp := make([][2]int, n+1)
	getDp(adj, visited, dp, weights, 1)

	visited = make([]bool, n+1)
	visited[1] = true
	res := make([]int, 0, n)

	isContain := dp[1][0] < dp[1][1]
	res = getRes(adj, visited, dp, weights, 1, isContain, res)

	_, _ = fmt.Fprintln(writer, integer.Max(dp[1][0], dp[1][1]))

	sort.Ints(res)
	for _, v := range res {
		_, _ = fmt.Fprintf(writer, "%d ", v)
	}
}

func getRes(adj [][]int, visited []bool, dp [][2]int, weights []int, cur int, isContain bool, res []int) []int {
	if isContain {
		res = append(res, cur)
	}

	for _, next := range adj[cur] {
		if visited[next] {
			continue
		}
		visited[next] = true

		if isContain {
			// cur 넣으면 next 무조건 뺌
			res = getRes(adj, visited, dp, weights, next, false, res)
		} else {
			// cur 빼면 next (빼는 것/넣는 것) 중 큰 것
			res = getRes(adj, visited, dp, weights, next, dp[next][0] < dp[next][1], res)
		}
	}
	return res
}

func getDp(adj [][]int, visited []bool, dp [][2]int, weights []int, cur int) {
	for _, next := range adj[cur] {
		if visited[next] {
			continue
		}
		visited[next] = true
		getDp(adj, visited, dp, weights, next)

		dp[cur][0] += integer.Max(dp[next][0], dp[next][1]) // cur 빼면 next (빼는 것/넣는 것) 중 큰 것
		dp[cur][1] += dp[next][0]                           // cur 넣으면 next 무조건 뺌
	}
	dp[cur][1] += weights[cur] // cur 넣기
}
