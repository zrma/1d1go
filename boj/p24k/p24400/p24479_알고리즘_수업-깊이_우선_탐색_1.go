package p24400

import (
	"io"
	"sort"
)

func Solve24479(reader io.Reader, writer io.Writer) {
	traverse := func(n, r int, graph [][]int, seq []int) {
		for i := 0; i <= n; i++ {
			sort.Ints(graph[i])
		}
		dfs(r, 0, graph, seq)
	}
	solveGraphTraversal(traverse, reader, writer)
}

func dfs(r, seqIdx int, graph [][]int, seq []int) int {
	seqIdx++
	seq[r] = seqIdx

	for _, v := range graph[r] {
		if seq[v] > 0 {
			continue
		}
		seqIdx = dfs(v, seqIdx, graph, seq)
	}
	return seqIdx
}
