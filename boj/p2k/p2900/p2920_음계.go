package p2900

import (
	"fmt"
	"io"
	"sort"
)

func Solve2920(reader io.Reader, writer io.Writer) {
	const n = 8
	arr := make([]int, n)

	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	if sort.SliceIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] }) {
		_, _ = fmt.Fprint(writer, "ascending")
		return
	}

	if sort.SliceIsSorted(arr, func(i, j int) bool { return arr[i] > arr[j] }) {
		_, _ = fmt.Fprint(writer, "descending")
	} else {
		_, _ = fmt.Fprint(writer, "mixed")
	}
}
