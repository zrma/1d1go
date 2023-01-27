package p11600

import (
	"fmt"
	"io"
	"sort"
)

func Solve11650(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([][2]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i][0], &arr[i][1])
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] < arr[j][0] {
			return true
		}
		if arr[i][0] > arr[j][0] {
			return false
		}
		return arr[i][1] < arr[j][1]
	})

	for _, v := range arr {
		_, _ = fmt.Fprintln(writer, v[0], v[1])
	}
}
