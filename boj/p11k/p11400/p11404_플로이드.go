package p11400

import (
	"fmt"
	"io"

	"1d1go/boj/p1k/p1300"
)

func Solve11404(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	const max = 1_000_000_000

	graph := make([][]int, n+1)
	for from := 1; from <= n; from++ {
		graph[from] = make([]int, n+1)
		for to := 1; to <= n; to++ {
			if from != to {
				graph[from][to] = max
			}
		}
	}

	for i := 0; i < m; i++ {
		var from, to, cost int
		_, _ = fmt.Fscan(reader, &from, &to, &cost)
		if cost < graph[from][to] {
			graph[from][to] = cost
		}
	}

	p1300.FloydWarshall(graph, n)

	for from := 1; from <= n; from++ {
		for to := 1; to <= n; to++ {
			if graph[from][to] >= max {
				_, _ = fmt.Fprint(writer, 0)
			} else {
				_, _ = fmt.Fprint(writer, graph[from][to])
			}

			if to != n {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}
}
