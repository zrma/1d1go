package p15600

import (
	"fmt"
)

func Solve15681(reader Reader, writer Writer) {
	var n, r, q int
	_, _ = fmt.Fscan(reader, &n, &r, &q)

	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	for i := 0; i < n-1; i++ {
		var u, v int
		_, _ = fmt.Fscan(reader, &u, &v)

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n+1)
	subTreeSize := make([]int, n+1)
	getSubTreeSize(adj, visited, subTreeSize, r)

	for i := 0; i < q; i++ {
		var u int
		_, _ = fmt.Fscan(reader, &u)

		_, _ = fmt.Fprintln(writer, subTreeSize[u])
	}
}

func getSubTreeSize(adj [][]int, visited []bool, size []int, r int) {
	visited[r] = true
	size[r] = 1

	for _, v := range adj[r] {
		if !visited[v] {
			getSubTreeSize(adj, visited, size, v)
			size[r] += size[v]
		}
	}
}
