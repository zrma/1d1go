package p11600

import (
	"fmt"
)

func Solve11659(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i+1])
		arr[i+1] += arr[i]
	}

	for i := 0; i < m; i++ {
		var begin, end int
		_, _ = fmt.Fscan(reader, &begin, &end)

		res := arr[end] - arr[begin-1]
		_, _ = fmt.Fprintln(writer, res)
	}
}
