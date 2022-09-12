package p1900

import (
	"fmt"

	"1d1go/utils/ds"
)

func Solve1927(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	h := ds.NewMinHeap(n)
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
