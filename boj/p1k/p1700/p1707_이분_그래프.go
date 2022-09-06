package p1700

import (
	"fmt"
)

func Solve1707(reader Reader, writer Writer) {
	var k int
	_, _ = fmt.Fscan(reader, &k)

	for i := 0; i < k; i++ {
		if isBipartiteGraph(reader) {
			_, _ = fmt.Fprintln(writer, "YES")
		} else {
			_, _ = fmt.Fprintln(writer, "NO")
		}
	}
}

func isBipartiteGraph(reader Reader) bool {
	var v, e int
	_, _ = fmt.Fscan(reader, &v, &e)

	graph := make([][]int, v+1)
	visited := make([]int, v+1)
	for i := 0; i < e; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := 1; i <= v; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			if !bfs(i, graph, visited) {
				return false
			}
		}
	}
	return true
}

func bfs(r int, graph [][]int, visited []int) bool {
	queue := make([]int, 0, len(visited))
	queueIdx := 0

	queue = append(queue, r)

	for queueIdx < len(queue) {
		here := queue[queueIdx]
		queueIdx++

		for _, v := range graph[here] {
			if visited[v] == 0 {
				visited[v] = -visited[here]
				queue = append(queue, v)
			} else if visited[v] == visited[here] {
				return false
			}
		}
	}
	return true
}
