package p1900

import (
	"fmt"
	"sort"
)

func Solve1920(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	var m int
	_, _ = fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		idx := sort.SearchInts(arr, v)
		if idx < len(arr) && arr[idx] == v {
			_, _ = fmt.Fprintln(writer, 1)
		} else {
			_, _ = fmt.Fprintln(writer, 0)
		}
	}
}
