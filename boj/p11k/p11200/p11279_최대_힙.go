package p11200

import (
	"fmt"
	"io"

	"1d1go/utils/ds"
)

func Solve11279(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	h := ds.NewMaxHeap(n)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if v == 0 {
			res, _ := h.Pop()
			_, _ = fmt.Fprintln(writer, res)
			continue
		}
		h.Push(v)
	}
}
