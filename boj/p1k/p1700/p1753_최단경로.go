package p1700

import (
	"container/heap"
	"fmt"
	"math"

	"1d1go/utils/ds"
)

func Solve1753(reader Reader, writer Writer) {
	var v, e, k int
	_, _ = fmt.Fscan(reader, &v, &e, &k)
	k -= 1

	graph := make([][]Edge, v)
	for i := 0; i < e; i++ {
		var u, v, w int
		_, _ = fmt.Fscan(reader, &u, &v, &w)
		u -= 1
		v -= 1
		graph[u] = append(graph[u], Edge{v, w})
	}

	dist := Dijkstra(v, e, k, graph)

	for i := 0; i < v; i++ {
		if dist[i] >= INF {
			_, _ = fmt.Fprintln(writer, "INF")
		} else {
			_, _ = fmt.Fprintln(writer, dist[i])
		}
	}
}

const INF = math.MaxInt32

func Dijkstra(v, e, k int, graph [][]Edge) []int {
	dist := make([]int, v)
	for i := 0; i < v; i++ {
		dist[i] = INF
	}
	dist[k] = 0

	pq := ds.NewPriorityQueue(e)
	heap.Push(pq, item{value: k, priority: 0})

	for pq.Len() > 0 {
		it := heap.Pop(pq).(item)
		u := it.value
		for _, e := range graph[u] {
			v := e.V
			w := e.W
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				heap.Push(pq, item{value: v, priority: dist[v]})
			}
		}
	}
	return dist
}

type Edge struct {
	V, W int
}

var _ ds.Priority = (*item)(nil)

type item struct {
	value    int
	priority int
}

func (i item) Priority() int { return i.priority }
