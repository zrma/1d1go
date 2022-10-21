package p4800

import (
	"fmt"
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
	parents := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parents[i] = -1
	}

	cycles := make([]int, 0, m)
	for i := 1; i <= m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		if find(parents, a) == find(parents, b) {
			cycles = append(cycles, a)
		} else {
			union(parents, a, b)
		}
	}

	visited := make([]bool, n+1)
	for _, cycle := range cycles {
		root := find(parents, cycle)
		visited[root] = true
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		root := find(parents, i)
		if visited[root] {
			continue
		}
		visited[root] = true
		cnt++
	}
	return cnt
}

func find(parents []int, x int) int {
	if parents[x] < 0 {
		return x
	}
	parents[x] = find(parents, parents[x])
	return parents[x]
}

func union(parents []int, a, b int) {
	a = find(parents, a)
	b = find(parents, b)
	if a == b {
		return
	}

	if parents[a] < parents[b] {
		parents[a] += parents[b]
		parents[b] = a
	} else {
		parents[b] += parents[a]
		parents[a] = b
	}
}

func Solve4803DFS(n, m int, reader Reader) int {
	graph := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

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
	graph := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

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
