package p24400

import (
	"sort"
)

func Solve24479(reader Reader, writer Writer) {
	traverse := func(n, r int) {
		for i := 0; i <= n; i++ {
			sort.Ints(graph[i])
		}
		dfs(r, 0)
	}
	solveGraphTraversal(traverse, reader, writer)
}

func dfs(r, seqIdx int) int {
	seqIdx++
	seq[r] = seqIdx

	for _, v := range graph[r] {
		if seq[v] > 0 {
			continue
		}
		seqIdx = dfs(v, seqIdx)
	}
	return seqIdx
}
