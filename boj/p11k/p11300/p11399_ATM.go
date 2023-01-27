package p11300

import (
	"fmt"
	"io"
	"sort"
)

func Solve11399(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	sort.Ints(arr)

	sum := 0
	for i, v := range arr {
		sum += v * (n - i)
	}
	_, _ = fmt.Fprint(writer, sum)
}
