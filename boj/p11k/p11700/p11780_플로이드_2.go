package p11700

import (
	"fmt"
	"io"
	"math"
)

func Solve11780(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	graph := make([][]node, n+1)
	for from := 1; from <= n; from++ {
		graph[from] = make([]node, n+1)
		for to := 1; to <= n; to++ {
			graph[from][to].prev = -1
			if from != to {
				graph[from][to].cost = math.MaxInt32
			}
		}
	}

	for i := 0; i < m; i++ {
		var from, to, cost int
		_, _ = fmt.Fscan(reader, &from, &to, &cost)
		if graph[from][to].cost > cost {
			graph[from][to].cost = cost
			graph[from][to].prev = from
		}
	}

	floydWarshall(graph, n)

	for from := 1; from <= n; from++ {
		for to := 1; to <= n; to++ {
			if graph[from][to].cost == math.MaxInt32 {
				_, _ = fmt.Fprint(writer, 0)
			} else {
				_, _ = fmt.Fprint(writer, graph[from][to].cost)
			}
			if to != n {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
	}

	for from := 1; from <= n; from++ {
		for to := 1; to <= n; to++ {
			if graph[from][to].prev == -1 {
				_, _ = fmt.Fprintln(writer, 0)
			} else {
				var path []int
				for cur := to; cur != -1; cur = graph[from][cur].prev {
					path = append(path, cur)
				}
				_, _ = fmt.Fprint(writer, len(path))
				for i := len(path) - 1; i >= 0; i-- {
					_, _ = fmt.Fprint(writer, " ", path[i])
				}
				_, _ = fmt.Fprintln(writer)
			}
		}
	}
}

func floydWarshall(graph [][]node, n int) {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if graph[i][k].cost+graph[k][j].cost < graph[i][j].cost {
					graph[i][j].cost = graph[i][k].cost + graph[k][j].cost
					graph[i][j].prev = graph[k][j].prev
				}
			}
		}
	}
}

type node struct {
	cost, prev int
}
