package p11700

import (
	"fmt"
	"io"
)

func Solve11725(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	graph := make([][]int, n+1)
	parents := make([]int, n+1)
	visited := make([]bool, n+1)

	for i := 0; i < n; i++ {
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	var queue []int
	queue = append(queue, 1)
	visited[1] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, v := range graph[p] {
			if !visited[v] {
				parents[v] = p
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	for i := 2; i <= n; i++ {
		_, _ = fmt.Fprintln(writer, parents[i])
	}
}
