package p24400

import (
	"fmt"
	"sort"
)

func Solve24480(reader Reader, writer Writer) {
	var n, m, r int
	_, _ = fmt.Fscan(reader, &n, &m, &r)

	graph = make([][]int, n+1)
	seq = make([]int, n+1)

	for i := 0; i < m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := 0; i <= n; i++ {
		sort.Slice(graph[i], func(j, k int) bool {
			return graph[i][j] > graph[i][k]
		})
	}

	dfs(r, 0)

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fprintln(writer, seq[i])
	}
}
