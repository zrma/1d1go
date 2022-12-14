package p15600

import (
	"fmt"
)

func Solve15650(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, 0, m)
	combinationDFS15650(arr, 0, n, m, writer)
}

func combinationDFS15650(arr []int, depth, n, m int, writer Writer) {
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

	for i := depth; i < n; i++ {
		arr = append(arr, i+1)
		combinationDFS15650(arr, i+1, n, m, writer)
		arr = arr[:len(arr)-1]
	}
}
