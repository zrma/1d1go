package p2900

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve2920(scanner Scanner, writer Writer) {
	const n = 8
	arr := make([]int, n)

	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
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
