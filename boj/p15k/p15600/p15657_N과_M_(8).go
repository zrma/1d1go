package p15600

import (
	"fmt"
	"sort"
)

func Solve15657(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Ints(arr)

	res := make([]int, m)

	combinationWithRepetitionDFS15657(arr, res, 0, 0, n, m, writer)
}

func combinationWithRepetitionDFS15657(arr []int, res []int, depth, start, n, m int, writer Writer) {
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

	for i := start; i < n; i++ {
		res[depth] = arr[i]
		combinationWithRepetitionDFS15657(arr, res, depth+1, i, n, m, writer)
	}
}
