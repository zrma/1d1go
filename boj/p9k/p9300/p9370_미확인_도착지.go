package p9300

import (
	"fmt"
	"io"
	"sort"

	"1d1go/boj/p1k/p1700"
)

func Solve9370(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		solve9370(reader, writer)
	}
}

func solve9370(reader io.Reader, writer io.Writer) {
	var n, m, t int
	_, _ = fmt.Fscan(reader, &n, &m, &t)

	var s, g, h int
	_, _ = fmt.Fscan(reader, &s, &g, &h)
	s -= 1
	g -= 1
	h -= 1

	graph := make([][]p1700.Edge, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]p1700.Edge, n)
	}

	const inf = 50_000 * 2_000

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = p1700.Edge{V: j, W: inf}
			graph[j][i] = p1700.Edge{V: i, W: inf}
		}
	}

	for i := 0; i < m; i++ {
		var a, b, d int
		_, _ = fmt.Fscan(reader, &a, &b, &d)
		a -= 1
		b -= 1
		d *= 2
		graph[a][b] = p1700.Edge{V: b, W: d}
		graph[b][a] = p1700.Edge{V: a, W: d}
	}
	graph[g][h] = p1700.Edge{V: graph[g][h].V, W: graph[g][h].W - 1}
	graph[h][g] = p1700.Edge{V: graph[h][g].V, W: graph[h][g].W - 1}

	dist := p1700.Dijkstra(n, m, s, graph)

	dest := make([]int, t)
	for i := 0; i < t; i++ {
		_, _ = fmt.Fscan(reader, &dest[i])
	}
	res := make([]int, 0, len(dest))
	for _, d := range dest {
		if dist[d-1]%2 == 1 {
			res = append(res, d)
		}
	}

	sort.Ints(res)

	for i, r := range res {
		if i == len(res)-1 {
			_, _ = fmt.Fprintln(writer, r)
		} else {
			_, _ = fmt.Fprint(writer, r, " ")
		}
	}
}
