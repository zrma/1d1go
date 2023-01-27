package p1100

import (
	"fmt"
	"io"
	"sort"

	"1d1go/utils/ds"
)

func Solve1197(reader io.Reader, writer io.Writer) {
	var v, e int
	_, _ = fmt.Fscan(reader, &v, &e)

	edges := make([]edge, e)
	for i := 0; i < e; i++ {
		var from, to, cost int
		_, _ = fmt.Fscan(reader, &from, &to, &cost)
		edges[i] = edge{from, to, cost}
	}

	res := kruskal(v, edges)
	_, _ = fmt.Fprint(writer, res)
}

func kruskal(v int, edges []edge) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	uf := ds.NewUnionFind(v)
	res := 0
	for _, edge := range edges {
		if uf.Find(edge.from) != uf.Find(edge.to) {
			uf.Union(edge.from, edge.to)
			res += edge.cost
		}
	}

	return res
}

type edge struct {
	from, to, cost int
}
