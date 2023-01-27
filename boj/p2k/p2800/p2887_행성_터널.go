package p2800

import (
	"fmt"
	"io"
	"sort"

	"1d1go/utils/ds"
)

func Solve2887(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	if n <= 1 {
		_, _ = fmt.Fprint(writer, 0)
		return
	}

	type planet struct {
		idx     int
		x, y, z int
	}

	planets := make([]planet, n)
	for i := 0; i < n; i++ {
		planets[i].idx = i
		_, _ = fmt.Fscan(reader, &planets[i].x, &planets[i].y, &planets[i].z)
	}

	edges := make([]edge, 0, 3*n-3)

	dist := func(from, to int) int {
		return to - from
	}

	sort.Slice(planets, func(i, j int) bool {
		return planets[i].x < planets[j].x
	})
	for i := 0; i < n-1; i++ {
		edges = append(edges, edge{planets[i].idx, planets[i+1].idx, dist(planets[i].x, planets[i+1].x)})
	}
	sort.Slice(planets, func(i, j int) bool {
		return planets[i].y < planets[j].y
	})
	for i := 0; i < n-1; i++ {
		edges = append(edges, edge{planets[i].idx, planets[i+1].idx, dist(planets[i].y, planets[i+1].y)})
	}
	sort.Slice(planets, func(i, j int) bool {
		return planets[i].z < planets[j].z
	})
	for i := 0; i < n-1; i++ {
		edges = append(edges, edge{planets[i].idx, planets[i+1].idx, dist(planets[i].z, planets[i+1].z)})
	}

	res := kruskal(n, edges)
	_, _ = fmt.Fprint(writer, res)
}

func kruskal(n int, edges []edge) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	uf := ds.NewUnionFind(n)
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
	from, to int
	cost     int
}
