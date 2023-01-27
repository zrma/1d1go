package p18800

import (
	"fmt"
	"io"
	"sort"
)

func Solve18870(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	exist := make(map[int]bool)
	set := make([]int, 0, n)
	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
		v := arr[i]
		if _, ok := exist[v]; !ok {
			set = append(set, v)
			exist[v] = true
		}
	}
	sort.Ints(set)

	for _, v := range arr {
		_, _ = fmt.Fprintf(writer, "%d ", sort.SearchInts(set, v))
	}
}
