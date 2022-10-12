package p11700

import (
	"container/heap"
	"fmt"

	"1d1go/utils/ds"
)

func Solve11779(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	graph := make([][]edge, n)
	for i := 0; i < m; i++ {
		var u, v, w int
		_, _ = fmt.Fscan(reader, &u, &v, &w)
		u -= 1
		v -= 1
		graph[u] = append(graph[u], edge{v, w})
	}

	var start, end int
	_, _ = fmt.Fscan(reader, &start, &end)
	start -= 1
	end -= 1

	dist := Dijkstra(n, m, start, graph)

	path := make([]int, 0, n)
	for i := end; i != start; i = dist[i].prev {
		path = append(path, i)
	}

	_, _ = fmt.Fprintln(writer, dist[end].dist)
	_, _ = fmt.Fprintln(writer, len(path)+1)
	_, _ = fmt.Fprintf(writer, "%d ", start+1)
	for i := len(path) - 1; i >= 0; i-- {
		_, _ = fmt.Fprintf(writer, "%d ", path[i]+1)
	}
}

func Dijkstra(n int, m int, start int, graph [][]edge) []dist {
	const inf = 1_000_000_000
	dist := make([]dist, n)
	for i := 0; i < n; i++ {
		dist[i].dist = inf
	}
	dist[start].dist = 0

	pq := ds.NewPriorityQueue(m)
	heap.Push(pq, item{value: start, priority: 0})

	for pq.Len() > 0 {
		it := heap.Pop(pq).(item)
		p := it.value

		if dist[p].dist < it.priority {
			continue
		}

		for _, e := range graph[p] {
			if dist[e.v].dist > dist[p].dist+e.w {
				dist[e.v].dist = dist[p].dist + e.w
				dist[e.v].prev = p
				heap.Push(pq, item{value: e.v, priority: dist[e.v].dist})
			}
		}
	}

	return dist
}

type edge struct {
	v, w int
}

type dist struct {
	dist, prev int
}

var _ ds.Priority = (*item)(nil)

type item struct {
	value    int
	priority int
}

func (i item) Priority() int { return i.priority }
