package p2100

import (
	"fmt"
	"io"

	"1d1go/boj/p17k/p17300"
	"1d1go/utils/ds"
)

func Solve2162(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var lines = make([][2]p17300.Point, n+1)
	for i := 1; i <= n; i++ {
		_, _ = fmt.Fscan(reader, &lines[i][0].X, &lines[i][0].Y, &lines[i][1].X, &lines[i][1].Y)
	}

	uf := ds.NewUnionFind(n)

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if _, ok := p17300.IsIntersect(lines[i][0], lines[i][1], lines[j][0], lines[j][1]); ok {
				uf.Union(i, j)
			}
		}
	}

	children := uf.Children()
	max := 0
	for _, child := range children {
		if child > max {
			max = child
		}
	}
	_, _ = fmt.Fprintln(writer, len(children))
	_, _ = fmt.Fprintln(writer, max)
}
