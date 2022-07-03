package p1900

import (
	"fmt"
	"sort"
)

func Solve1931(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr2D := make([][2]int, n)
	for i := range arr2D {
		_, _ = fmt.Fscan(reader, &arr2D[i][0], &arr2D[i][1])
	}

	sort.Slice(arr2D, func(i, j int) bool {
		if arr2D[i][1] < arr2D[j][1] {
			return true
		}
		return arr2D[i][1] == arr2D[j][1] && arr2D[i][0] < arr2D[j][0]
	})

	cur := arr2D[0][1]
	cnt := 1
	for i := 1; i < n; i++ {
		if cur > arr2D[i][0] {
			continue
		}
		cur = arr2D[i][1]
		cnt++
	}
	_, _ = fmt.Fprint(writer, cnt)
}
