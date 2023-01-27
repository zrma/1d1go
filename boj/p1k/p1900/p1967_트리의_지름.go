package p1900

import (
	"fmt"
	"io"

	"1d1go/boj/p1k/p1100"
)

func Solve1967(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	tree := make([]p1100.Node, n+1)
	for i := 1; i <= n; i++ {
		var from, to, cost int
		_, _ = fmt.Fscan(reader, &from, &to, &cost)

		tree[from].Children = append(tree[from].Children, p1100.Edge{Idx: to, Cost: cost})
		tree[to].Children = append(tree[to].Children, p1100.Edge{Idx: from, Cost: cost})
	}

	farthestNode0, _ := p1100.GetFarthestNode(tree, 1)
	_, dist := p1100.GetFarthestNode(tree, farthestNode0)
	_, _ = fmt.Fprint(writer, dist)
}
