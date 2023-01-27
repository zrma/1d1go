package p2200

import (
	"fmt"
	"io"
	"sort"
)

func Solve2212(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	distances := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		distances[i] = arr[i+1] - arr[i]
	}
	sort.Ints(distances)

	res := 0
	for i := 0; i < n-k; i++ {
		res += distances[i]
	}
	_, _ = fmt.Fprintf(writer, "%d", res)
}
