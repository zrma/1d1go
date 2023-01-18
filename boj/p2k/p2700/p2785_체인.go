package p2700

import (
	"fmt"
	"sort"
)

func Solve2785(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	res := 0
	for len(arr) > 1 {
		res++
		arr[len(arr)-2] += arr[len(arr)-1]
		arr = arr[:len(arr)-1]

		arr[0]--
		if arr[0] == 0 {
			arr = arr[1:]
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
