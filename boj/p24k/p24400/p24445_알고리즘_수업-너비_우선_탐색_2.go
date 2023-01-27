package p24400

import (
	"io"
	"sort"
)

func Solve24445(reader io.Reader, writer io.Writer) {
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
