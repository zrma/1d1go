package p1200

import (
	"fmt"
	"sort"
)

func Solve1260(reader Reader, writer Writer) {
	var n, m, r int
	_, _ = fmt.Fscan(reader, &n, &m, &r)

	graph := make([][]int, n+1)
	visited := make([]bool, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := 0; i <= n; i++ {
		sort.Ints(graph[i])
	}

	dfs(r, writer, graph, visited)
	_, _ = fmt.Fprintln(writer)

	visited = make([]bool, n+1)
	bfs(r, writer, graph, visited)
}

func dfs(r int, writer Writer, graph [][]int, visited []bool) {
	visited[r] = true
	_, _ = fmt.Fprint(writer, r)

	for _, v := range graph[r] {
		if visited[v] {
			continue
		}
		_, _ = fmt.Fprint(writer, " ")
		dfs(v, writer, graph, visited)
	}
}

func bfs(r int, writer Writer, graph [][]int, visited []bool) {
	queue := make([]int, 0, len(visited))
	queueIdx := 0

	queue = append(queue, r)
	visited[r] = true
	_, _ = fmt.Fprint(writer, r)

	for queueIdx < len(queue) {
		here := queue[queueIdx]
		queueIdx++

		for _, next := range graph[here] {
			if visited[next] {
				continue
			}
			visited[next] = true
			_, _ = fmt.Fprint(writer, " ", next)
			queue = append(queue, next)
		}
	}
}
