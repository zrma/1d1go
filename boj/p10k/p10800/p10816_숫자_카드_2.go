package p10800

import (
	"fmt"
	"io"
	"sort"
)

func Solve10816(reader io.Reader, writer io.Writer) {
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

		next := v + 1

		idx := sort.SearchInts(arr, v)
		if idx >= len(arr) || arr[idx] != v {
			_, _ = fmt.Fprint(writer, 0, " ")
			continue
		}

		nextIdx := sort.SearchInts(arr, next)
		_, _ = fmt.Fprint(writer, nextIdx-idx, " ")
	}
}

func Solve10816WithMap(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		counts[v]++
	}

	var m int
	_, _ = fmt.Fscan(reader, &m)

	for i := 0; i < m; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		_, _ = fmt.Fprint(writer, counts[v], " ")
	}
}
