package p1900

import (
	"fmt"
)

func Solve1904(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	maxLen := n + 2
	arr := make([]int64, maxLen)

	const m = 15746

	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
		if arr[i] >= m {
			arr[i] %= m
		}
	}
	_, _ = fmt.Fprint(writer, arr[n])
}
