package p15600

import (
	"fmt"
	"io"
)

func Solve15649(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	visited := make([]bool, n)
	arr := make([]int, m)

	permutationDFS15649(arr, visited, 0, n, m, writer)
}

func permutationDFS15649(arr []int, visited []bool, depth, n, m int, writer io.Writer) {
	if depth == m {
		for i, v := range arr {
			_, _ = fmt.Fprint(writer, v)
			if i != m-1 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
		return
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		arr[depth] = i + 1
		permutationDFS15649(arr, visited, depth+1, n, m, writer)
		visited[i] = false
	}
}
