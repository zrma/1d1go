package p15600

import (
	"fmt"
)

func Solve15652(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, 0, m)
	combinationWithRepetitionDFS15652(arr, 1, n, m, writer)
}

func combinationWithRepetitionDFS15652(arr []int, depth, n, m int, writer Writer) {
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
		combinationWithRepetitionDFS15652(arr, i+1, n, m, writer)
		arr = arr[:len(arr)-1]
	}
}
