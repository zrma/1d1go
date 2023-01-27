package p11700

import (
	"fmt"
	"io"
)

func Solve11724(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		var u, v int
		_, _ = fmt.Fscan(reader, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n+1)
	var count int
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i, graph, visited)
			count++
		}
	}
	_, _ = fmt.Fprint(writer, count)
}

func dfs(r int, graph [][]int, visited []bool) {
	visited[r] = true
	for _, visit := range graph[r] {
		if !visited[visit] {
			dfs(visit, graph, visited)
		}
	}
}
