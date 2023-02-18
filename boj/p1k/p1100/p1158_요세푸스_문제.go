package p1100

import (
	"fmt"
	"io"
)

func Solve1158(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	var ans []int
	idx := 0
	for len(arr) > 0 {
		idx = (idx + k - 1) % len(arr)
		ans = append(ans, arr[idx])
		arr = append(arr[:idx], arr[idx+1:]...)
	}

	_, _ = fmt.Fprintf(writer, "<%d", ans[0])
	for _, v := range ans[1:] {
		_, _ = fmt.Fprintf(writer, ", %d", v)
	}
	_, _ = fmt.Fprint(writer, ">")
}
