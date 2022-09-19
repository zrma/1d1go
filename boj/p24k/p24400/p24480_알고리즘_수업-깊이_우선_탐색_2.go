package p24400

import (
	"sort"
)

func Solve24480(reader Reader, writer Writer) {
	traverse := func(n, r int, graph [][]int, seq []int) {
		for i := 0; i <= n; i++ {
			sort.Slice(graph[i], func(j, k int) bool {
				return graph[i][j] > graph[i][k]
			})
		}
		dfs(r, 0, graph, seq)
	}
	solveGraphTraversal(traverse, reader, writer)
}
