package p1500

import (
	"fmt"

	"1d1go/boj/p1k/p1700"
	"1d1go/utils/integer"
)

func Solve1504(reader Reader, writer Writer) {
	var n, e int
	_, _ = fmt.Fscan(reader, &n, &e)

	graph := make([][]p1700.Edge, n)
	for i := 0; i < e; i++ {
		var a, b, c int
		_, _ = fmt.Fscan(reader, &a, &b, &c)
		a -= 1
		b -= 1
		graph[a] = append(graph[a], p1700.Edge{V: b, W: c})
		graph[b] = append(graph[b], p1700.Edge{V: a, W: c})
	}

	var v1, v2 int
	_, _ = fmt.Fscan(reader, &v1, &v2)
	v1 -= 1
	v2 -= 1
	dist := p1700.Dijkstra(n, e, 0, graph)
	dist1 := p1700.Dijkstra(n, e, v1, graph)
	dist2 := p1700.Dijkstra(n, e, v2, graph)

	res := integer.Min(dist[v1]+dist1[v2]+dist2[n-1], dist[v2]+dist2[v1]+dist1[n-1])
	if res >= p1700.INF {
		_, _ = fmt.Fprint(writer, -1)
	} else {
		_, _ = fmt.Fprint(writer, res)
	}
}
