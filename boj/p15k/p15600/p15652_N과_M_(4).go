package p15600

import (
	"fmt"
	"strconv"
)

func Solve15652(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, 0, m)
	combinationWithRepetitionDFS(arr, 1, n, m, writer)
}

func combinationWithRepetitionDFS(arr []int, depth, n, m int, writer Writer) {
	if len(arr) == m {
		for i, v := range arr {
			_, _ = fmt.Fprint(writer, v)
			if i != m-1 {
				_, _ = fmt.Fprint(writer, " ")
			}
		}
		_, _ = fmt.Fprintln(writer)
		return
	}

	for i := depth - 1; i < n; i++ {
		arr = append(arr, i+1)
		combinationWithRepetitionDFS(arr, i+1, n, m, writer)
		arr = arr[:len(arr)-1]
	}
}
