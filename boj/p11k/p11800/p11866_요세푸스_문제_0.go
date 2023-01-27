package p11800

import (
	"fmt"
	"io"
)

func Solve11866(reader io.Reader, writer io.Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}

	_, _ = fmt.Fprint(writer, "<")
	for len(arr) > 1 {
		mod := (len(arr) + k - 1) % len(arr)
		_, _ = fmt.Fprintf(writer, "%d, ", arr[mod])
		arr = append(arr[mod+1:], arr[:mod]...)
	}
	_, _ = fmt.Fprint(writer, arr[0], ">")
}
