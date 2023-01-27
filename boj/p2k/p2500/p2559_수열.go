package p2500

import (
	"fmt"
	"io"
)

func Solve2559(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	sum := 0
	for i := 0; i < k; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
		sum += arr[i]
	}
	max := sum
	for i := k; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])

		sum -= arr[i-k]
		sum += arr[i]
		if sum > max {
			max = sum
		}
	}
	_, _ = fmt.Fprint(writer, max)
}
