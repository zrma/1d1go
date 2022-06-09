package p15600

import (
	"fmt"
	"strconv"
)

func Solve15649(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	visited := make([]bool, n)
	arr := make([]int, m)

	permutationDFS(arr, visited, 0, n, m, writer)
}

func permutationDFS(arr []int, visited []bool, depth, n, m int, writer Writer) {
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
		permutationDFS(arr, visited, depth+1, n, m, writer)
		visited[i] = false
	}
}
