package p11700

import (
	"fmt"
)

func Solve11724(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	graph = make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		var u, v int
		_, _ = fmt.Fscan(reader, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited = make([]bool, n+1)
	var count int
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}
	_, _ = fmt.Fprint(writer, count)
}

var (
	graph   [][]int
	visited []bool
)

func dfs(r int) {
	visited[r] = true
	for _, visit := range graph[r] {
		if !visited[visit] {
			dfs(visit)
		}
	}
}
