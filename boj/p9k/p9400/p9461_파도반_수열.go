package p9400

import (
	"fmt"
	"io"
)

func Solve9461(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	const maxLen = 100 + 1
	arr := make([]int, maxLen)

	arr[0] = 0
	arr[1] = 1
	arr[2] = 1
	arr[3] = 1
	arr[4] = 2

	for i := 5; i < maxLen; i++ {
		arr[i] = arr[i-1] + arr[i-5]
	}

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		res := arr[v]
		_, _ = fmt.Fprintln(writer, res)
	}
}
