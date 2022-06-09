package p15600

import (
	"fmt"
	"strconv"
)

func Solve15651(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, m)
	permutationWithRepetitionDFS(arr, 0, n, m, writer)
}

func permutationWithRepetitionDFS(arr []int, depth, n, m int, writer Writer) {
	if depth >= m {
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
		arr[depth] = i + 1
		permutationWithRepetitionDFS(arr, depth+1, n, m, writer)
	}
}
