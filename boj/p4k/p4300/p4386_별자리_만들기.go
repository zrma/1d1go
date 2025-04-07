package p4300

import (
	"fmt"
	"io"
	"math"
	"sort"

	"1d1go/utils/ds"
)

func Solve4386(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	pos := make([][2]float64, n)
	for i := range pos {
		_, _ = fmt.Fscan(reader, &pos[i][0], &pos[i][1])
	}

	edges := make([]edge, 0, n*n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, dist(pos[i], pos[j])})
		}
	}

	res := kruskal(n, edges)
	_, _ = fmt.Fprintf(writer, "%.2f", res)
}

func kruskal(n int, edges []edge) float64 {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	uf := ds.NewUnionFind(n)
	res := 0.0
	for _, edge := range edges {
		if uf.Find(edge.from) != uf.Find(edge.to) {
			uf.Union(edge.from, edge.to)
			res += edge.cost
		}
	}

	return res
}

type edge struct {
	from, to int
	cost     float64
}

func dist(pos0 [2]float64, pos1 [2]float64) float64 {
	return math.Sqrt((pos0[0]-pos1[0])*(pos0[0]-pos1[0]) + (pos0[1]-pos1[1])*(pos0[1]-pos1[1]))
}
