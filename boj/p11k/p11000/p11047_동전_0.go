package p11000

import (
	"fmt"
	"io"
)

func Solve11047(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	cnt := 0
	idx := 0
	for k > 0 && idx < n {
		cur := arr[idx]
		if k < cur {
			idx++
			continue
		}
		v := k / cur
		k -= v * cur
		cnt += v
	}
	_, _ = fmt.Fprint(writer, cnt)
}
