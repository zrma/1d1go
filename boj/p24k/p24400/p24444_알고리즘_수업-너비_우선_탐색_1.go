package p24400

import (
	"fmt"
	"sort"
)

func Solve24444(reader Reader, writer Writer) {
	traverse := func(n, r int) {
		for i := 0; i <= n; i++ {
			sort.Ints(graph[i])
		}
		bfs(r)
	}
	solveGraphTraversal(traverse, reader, writer)
}

type traversalFunc func(n, r int)

func solveGraphTraversal(traverse traversalFunc, reader Reader, writer Writer) {
	var n, m, r int
	_, _ = fmt.Fscan(reader, &n, &m, &r)

	graph = make([][]int, n+1)
	seq = make([]int, n+1)

	for i := 0; i < m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	traverse(n, r)

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fprintln(writer, seq[i])
	}
}

var (
	graph [][]int
	seq   []int
)

func bfs(r int) {
	seqIdx := 1
	seq[r] = seqIdx

	queue := make([]int, 0, len(seq))
	queueIdx := 0

	queue = append(queue, r)
	for queueIdx < len(queue) {
		here := queue[queueIdx]
		queueIdx++

		for _, visit := range graph[here] {
			if seq[visit] > 0 {
				continue
			}
			seqIdx++
			seq[visit] = seqIdx
			queue = append(queue, visit)
		}
	}
}
