package p2600

import (
	"fmt"
)

func Solve2606(reader Reader, writer Writer) {
	var computers, pairs int
	_, _ = fmt.Fscan(reader, &computers, &pairs)

	graph := make([][]int, computers+1)
	visited := make([]bool, computers+1)
	for i := 0; i < pairs; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	queue := make([]int, 0, pairs)
	queueIdx := 0
	count := 0

	queue = append(queue, 1)
	visited[1] = true

	for queueIdx < len(queue) {
		here := queue[queueIdx]
		queueIdx++

		for _, next := range graph[here] {
			if visited[next] {
				continue
			}
			count++
			visited[next] = true
			queue = append(queue, next)
		}
	}
	_, _ = fmt.Fprint(writer, count)
}
