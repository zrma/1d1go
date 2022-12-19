package p6400

import (
	"fmt"
	"sort"

	"1d1go/utils/ds"
)

func Solve6497(reader Reader, writer Writer) {
	for {
		var m, n int
		_, _ = fmt.Fscan(reader, &m, &n)

		if m == 0 && n == 0 {
			break
		}

		uf := ds.NewUnionFind(m)
		type edge struct {
			from, to, cost int
		}
		edges := make([]edge, n)
		for i := range edges {
			_, _ = fmt.Fscan(reader, &edges[i].from, &edges[i].to, &edges[i].cost)
		}
		sort.Slice(edges, func(i, j int) bool {
			return edges[i].cost < edges[j].cost
		})

		tot := 0
		res := 0
		for _, edge := range edges {
			tot += edge.cost
			if uf.Find(edge.from) != uf.Find(edge.to) {
				uf.Union(edge.from, edge.to)
				res += edge.cost
			}
		}
		_, _ = fmt.Fprintln(writer, tot-res)
	}
}
