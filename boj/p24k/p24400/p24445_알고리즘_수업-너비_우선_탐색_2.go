package p24400

import (
	"sort"
)

func Solve24445(reader Reader, writer Writer) {
	traverse := func(n, r int, graph [][]int, seq []int) {
		for i := 0; i <= n; i++ {
			sort.Slice(graph[i], func(j, k int) bool {
				return graph[i][j] > graph[i][k]
			})
		}
		bfs(r, graph, seq)
	}
	solveGraphTraversal(traverse, reader, writer)
}
