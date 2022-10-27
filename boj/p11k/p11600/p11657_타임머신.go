package p11600

import (
	"fmt"
)

func Solve11657(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	edges := make([]edge, m)
	for i := range edges {
		_, _ = fmt.Fscan(reader, &edges[i].u, &edges[i].v, &edges[i].w)
	}

	dist, hasCycle := bellmanFord(n, edges)
	if hasCycle {
		_, _ = fmt.Fprintln(writer, -1)
		return
	}

	for i := 2; i <= n; i++ {
		if dist[i] == inf {
			_, _ = fmt.Fprintln(writer, -1)
		} else {
			_, _ = fmt.Fprintln(writer, dist[i])
		}
	}
}

type edge struct {
	u, v, w int
}

const inf = 1_000_000_000

func bellmanFord(n int, edges []edge) ([]int, bool) {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[1] = 0

	for i := 0; i < n; i++ {
		for _, e := range edges {
			from := e.u
			to := e.v
			cost := e.w
			if dist[from] != inf && dist[to] > dist[from]+cost {
				dist[to] = dist[from] + cost

				if i == n-1 {
					return dist, true
				}
			}
		}
	}

	return dist, false
}
