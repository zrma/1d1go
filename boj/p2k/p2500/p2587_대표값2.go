package p2500

import (
	"fmt"
	"io"
	"sort"
)

func Solve2587(reader io.Reader, writer io.Writer) {
	const size = 5
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Ints(arr)

	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}

	_, _ = fmt.Fprintln(writer, int(sum/float64(size)))
	_, _ = fmt.Fprintln(writer, arr[size/2])
}
