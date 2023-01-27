package p1800

import (
	"fmt"
	"io"
)

func Solve1865(reader io.Reader, writer io.Writer) {
	var tc int
	_, _ = fmt.Fscan(reader, &tc)

	for ; tc > 0; tc-- {
		var n, m, w int
		_, _ = fmt.Fscan(reader, &n, &m, &w)

		edges := make([]edge, 0, 2*m+w)
		for i := 0; i < m; i++ {
			var u, v, c int64
			_, _ = fmt.Fscan(reader, &u, &v, &c)
			edges = append(edges, edge{u: u, v: v, w: c})
			edges = append(edges, edge{u: v, v: u, w: c})
		}

		for i := 0; i < w; i++ {
			var u, v, c int64
			_, _ = fmt.Fscan(reader, &u, &v, &c)
			edges = append(edges, edge{u: u, v: v, w: -c})
		}

		if hasCycle(n, edges) {
			_, _ = fmt.Fprintln(writer, "YES")
		} else {
			_, _ = fmt.Fprintln(writer, "NO")
		}
	}
}

type edge struct {
	u, v, w int64
}

const inf = 10_001

func hasCycle(n int, edges []edge) bool {
	dist := make([]int64, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[1] = 0

	for i := 0; i < n; i++ {
		for _, e := range edges {
			from := e.u
			to := e.v
			cost := e.w
			if dist[to] > dist[from]+cost {
				dist[to] = dist[from] + cost

				if i == n-1 {
					return true
				}
			}
		}
	}
	return false
}
