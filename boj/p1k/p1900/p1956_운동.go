package p1900

import (
	"fmt"
	"math"

	"1d1go/boj/p1k/p1300"
)

func Solve1956(reader Reader, writer Writer) {
	var v, e int
	_, _ = fmt.Fscan(reader, &v, &e)

	graph := make([][]int, v+1)
	for i := 1; i <= v; i++ {
		graph[i] = make([]int, v+1)
		for j := 1; j <= v; j++ {
			graph[i][j] = math.MaxInt32
		}
	}

	for i := 0; i < e; i++ {
		var from, to, weight int
		_, _ = fmt.Fscan(reader, &from, &to, &weight)
		graph[from][to] = weight
	}

	p1300.FloydWarshall(graph, v)

	var min = math.MaxInt32
	for i := 1; i <= v; i++ {
		if graph[i][i] < min {
			min = graph[i][i]
		}
	}

	if min == math.MaxInt32 {
		_, _ = fmt.Fprint(writer, -1)
	} else {
		_, _ = fmt.Fprint(writer, min)
	}
}
