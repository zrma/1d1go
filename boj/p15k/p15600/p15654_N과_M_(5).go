package p15600

import (
	"fmt"
	"io"
	"sort"
)

func Solve15654(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Ints(arr)

	visited := make([]bool, n)
	res := make([]int, m)

	permutationDFS15654(arr, visited, res, 0, n, m, writer)
}

func permutationDFS15654(arr []int, visited []bool, res []int, depth, n, m int, writer io.Writer) {
	if depth == m {
		for i, v := range res {
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
		res[depth] = arr[i]
		permutationDFS15654(arr, visited, res, depth+1, n, m, writer)
		visited[i] = false
	}
}
