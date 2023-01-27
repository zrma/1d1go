package p1900

import (
	"fmt"
	"io"

	"1d1go/utils/ds"
)

type solve1976Func func(int, int, io.Reader) bool

func Solve1976(reader io.Reader, writer io.Writer, f solve1976Func) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	if f(n, m, reader) {
		_, _ = fmt.Fprint(writer, "YES")
	} else {
		_, _ = fmt.Fprint(writer, "NO")
	}
}

func Solve1976UnionFind(n, m int, reader io.Reader) bool {
	uf := ds.NewUnionFind(n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var x int
			_, _ = fmt.Fscan(reader, &x)

			if i <= j {
				continue
			}

			if x == 0 {
				continue
			}

			uf.Union(i, j)
		}
	}

	var res int
	for i := 0; i < m; i++ {
		var a int
		_, _ = fmt.Fscan(reader, &a)

		if i == 0 {
			res = uf.Find(a)
		}
		if res != uf.Find(a) {
			return false
		}
	}

	return true
}

func Solve1976DFS(n, m int, reader io.Reader) bool {
	graph := scanGraph(n, reader)

	var start int
	_, _ = fmt.Fscan(reader, &start)

	visited := make([]bool, n+1)
	dfs1976(graph, visited, start)

	for i := 1; i < m; i++ {
		var a int
		_, _ = fmt.Fscan(reader, &a)

		if !visited[a] {
			return false
		}
	}
	return true
}

func dfs1976(graph [][]int, visited []bool, start int) {
	visited[start] = true

	for _, next := range graph[start] {
		if visited[next] {
			continue
		}

		dfs1976(graph, visited, next)
	}
}

func Solve1976BFS(n, m int, reader io.Reader) bool {
	graph := scanGraph(n, reader)

	var start int
	_, _ = fmt.Fscan(reader, &start)

	visited := make([]bool, n+1)
	bfs1976(graph, visited, start)

	for i := 1; i < m; i++ {
		var a int
		_, _ = fmt.Fscan(reader, &a)

		if !visited[a] {
			return false
		}
	}
	return true
}

func bfs1976(graph [][]int, visited []bool, start int) {
	queue := make([]int, 0, len(graph))
	queue = append(queue, start)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		visited[cur] = true

		for _, next := range graph[cur] {
			if visited[next] {
				continue
			}

			queue = append(queue, next)
		}
	}
}

func scanGraph(n int, reader io.Reader) [][]int {
	graph := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0, n)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var x int
			_, _ = fmt.Fscan(reader, &x)

			if i <= j {
				continue
			}

			if x == 0 {
				continue
			}

			graph[i] = append(graph[i], j)
			graph[j] = append(graph[j], i)
		}
	}
	return graph
}
