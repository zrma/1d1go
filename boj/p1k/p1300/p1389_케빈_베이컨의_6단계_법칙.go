package p1300

import (
	"fmt"
	"io"
	"math"
)

func Solve1389(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	graph := make([][]int, n+1)
	for from := 1; from <= n; from++ {
		graph[from] = make([]int, n+1)
		for to := 1; to <= n; to++ {
			if from != to {
				graph[from][to] = math.MaxInt32
			}
		}
	}

	for i := 0; i < m; i++ {
		var from, to int
		_, _ = fmt.Fscan(reader, &from, &to)
		graph[from][to] = 1
		graph[to][from] = 1
	}

	FloydWarshall(graph, n)

	var min = math.MaxInt32
	var res = 1
	for from := 1; from <= n; from++ {
		var sum = 0
		for to := 1; to <= n; to++ {
			sum += graph[from][to]
		}
		if sum < min {
			min = sum
			res = from
		}
	}

	_, _ = fmt.Fprint(writer, res)
}

func FloydWarshall(graph [][]int, n int) {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if graph[i][k]+graph[k][j] < graph[i][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}
}

func Solve1389WithBFS(reader io.Reader, writer io.Writer) {
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

	res := 0
	min := math.MaxInt32
	for from := 1; from <= n; from++ {
		cnt := 0
		for to := 1; to <= n; to++ {
			if from == to {
				continue
			}
			visited := make([]bool, n+1)
			cnt += bfs(graph, from, to, n, visited)
		}
		if cnt < min {
			min = cnt
			res = from
		}
	}
	_, _ = fmt.Fprint(writer, res)
}

func bfs(graph [][]int, from, to, n int, visited []bool) int {
	queue := make([][2]int, 0, n)
	queue = append(queue, [2]int{from, 0})

	for len(queue) > 0 {
		here, cnt := queue[0][0], queue[0][1]
		queue = queue[1:]
		if here == to {
			return cnt
		}
		visited[here] = true
		for _, v := range graph[here] {
			if !visited[v] {
				queue = append(queue, [2]int{v, cnt + 1})
			}
		}
	}
	return 0
}
