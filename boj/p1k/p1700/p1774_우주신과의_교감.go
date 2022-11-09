package p1700

import (
	"fmt"
	"math"
	"sort"

	"1d1go/utils/ds"
)

func Solve1774(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	uf := ds.NewUnionFind(n)
	pos := make([][2]int, n)
	for i := range pos {
		_, _ = fmt.Fscan(reader, &pos[i][0], &pos[i][1])
	}

	type edge struct {
		from, to int
		cost     float64
	}

	const max float64 = 1_000_000_000
	edges := make([]edge, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist := distance(pos[i], pos[j])
			if i == j {
				dist = max
			}
			edges[i*n+j] = edge{i, j, dist}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	for i := 0; i < m; i++ {
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)
		uf.Union(from-1, to-1)
	}

	res := 0.0
	for _, edge := range edges {
		if uf.Find(edge.from) != uf.Find(edge.to) {
			uf.Union(edge.from, edge.to)
			res += edge.cost
		}
	}

	_, _ = fmt.Fprintf(writer, "%.2f", res)
}

func distance(pos0 [2]int, pos1 [2]int) float64 {
	return math.Sqrt(math.Pow(float64(pos0[0]-pos1[0]), 2) + math.Pow(float64(pos0[1]-pos1[1]), 2))
}
