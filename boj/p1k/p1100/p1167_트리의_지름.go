package p1100

import (
	"fmt"
	"io"
)

func Solve1167(reader io.Reader, writer io.Writer) {
	var v int
	_, _ = fmt.Fscan(reader, &v)

	tree := make([]Node, v+1)
	for i := 1; i <= v; i++ {
		var idx int
		_, _ = fmt.Fscan(reader, &idx)

		for {
			var n int
			_, _ = fmt.Fscan(reader, &n)
			if n == -1 {
				break
			}
			var c int
			_, _ = fmt.Fscan(reader, &c)
			tree[idx].Children = append(tree[idx].Children, Edge{Idx: n, Cost: c})
		}
	}

	farthestNode0, _ := GetFarthestNode(tree, 1)
	_, dist := GetFarthestNode(tree, farthestNode0)
	_, _ = fmt.Fprint(writer, dist)

}

func GetFarthestNode(tree []Node, start int) (farthestNode int, maxCost int) {
	visited := make([]bool, len(tree))
	visited[start] = true

	costs := make([]int, len(tree))

	queue := []int{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, child := range tree[cur].Children {
			if visited[child.Idx] {
				continue
			}

			visited[child.Idx] = true
			queue = append(queue, child.Idx)
			costs[child.Idx] = costs[cur] + child.Cost
			if maxCost < costs[child.Idx] {
				maxCost = costs[child.Idx]
				farthestNode = child.Idx
			}
		}
	}
	return
}

type Node struct {
	Children []Edge
}

type Edge struct {
	Idx  int
	Cost int
}
