package p4800

import (
	"fmt"

	"1d1go/utils/ds"
)

type solve4803Func func(int, int, Reader) int

func Solve4803(reader Reader, writer Writer, f solve4803Func) {
	var n, m int

	for cnt := 1; ; cnt++ {
		_, _ = fmt.Fscan(reader, &n, &m)
		if n == 0 && m == 0 {
			break
		}

		res := f(n, m, reader)
		switch res {
		case 0:
			_, _ = fmt.Fprintf(writer, "Case %d: No trees.\n", cnt)
		case 1:
			_, _ = fmt.Fprintf(writer, "Case %d: There is one tree.\n", cnt)
		default:
			_, _ = fmt.Fprintf(writer, "Case %d: A forest of %d trees.\n", cnt, res)
		}
	}
}

func Solve4803UnionFind(n, m int, reader Reader) int {
	uf := ds.NewUnionFind(n)

	cycles := make([]int, 0, m)
	for i := 1; i <= m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		if uf.Find(a) == uf.Find(b) {
			cycles = append(cycles, a)
		} else {
			uf.Union(a, b)
		}
	}

	visited := make([]bool, n+1)
	for _, cycle := range cycles {
		root := uf.Find(cycle)
		visited[root] = true
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		root := uf.Find(i)
		if visited[root] {
			continue
		}
		visited[root] = true
		cnt++
	}
	return cnt
}

func Solve4803DFS(n, m int, reader Reader) int {
	graph := scanGraph(n, m, reader)

	visited := make([]bool, n+1)
	cnt := 0
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}

		if dfs4803(graph, visited, i, 0) {
			cnt++
		}
	}
	return cnt
}

func dfs4803(graph [][]int, visited []bool, before, curr int) bool {
	if visited[before] {
		return false
	}
	visited[before] = true

	for _, next := range graph[before] {
		if next == curr {
			continue
		}
		if !dfs4803(graph, visited, next, before) {
			return false
		}
	}
	return true
}

func Solve4803BFS(n, m int, reader Reader) int {
	graph := scanGraph(n, m, reader)

	visited := make([]bool, n+1)
	cnt := 0
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}

		if bfs4803(graph, visited, i) {
			cnt++
		}
	}
	return cnt
}

func bfs4803(graph [][]int, visited []bool, i int) bool {
	queue := make([]int, 0, len(graph))
	queue = append(queue, i)

	nodeCnt := 0
	edgeCnt := 0

	for len(queue) > 0 {
		nodeCnt++
		p := queue[0]
		queue = queue[1:]
		visited[p] = true

		for _, to := range graph[p] {
			edgeCnt++
			if visited[to] {
				continue
			}
			queue = append(queue, to)
		}
	}
	return nodeCnt-1 == edgeCnt/2
}

func scanGraph(n, m int, reader Reader) [][]int {
	graph := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}
