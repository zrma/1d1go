package p25300

import (
	"fmt"
	"io"
	"sort"
)

func Solve25305(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	_, _ = fmt.Fprint(writer, arr[k-1])
}
