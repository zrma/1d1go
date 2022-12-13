package p1200

import (
	"fmt"

	"1d1go/boj/p1k/p1700"
)

func Solve1238(reader Reader, writer Writer) {
	var n, m, x int
	_, _ = fmt.Fscan(reader, &n, &m, &x)

	graph := make([][]p1700.Edge, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		_, _ = fmt.Fscan(reader, &u, &v, &w)
		u -= 1
		v -= 1
		graph[u] = append(graph[u], p1700.Edge{V: v, W: w})
	}

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = p1700.Dijkstra(n, m, i, graph)
	}

	ans := 0
	for i := 0; i < n; i++ {
		if cost := d[i][x-1] + d[x-1][i]; cost > ans {
			ans = cost
		}
	}

	_, _ = fmt.Fprint(writer, ans)
}
