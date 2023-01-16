package p1700

import (
	"fmt"

	"1d1go/utils/ds"
)

func Solve1715(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	if n == 1 {
		_, _ = fmt.Fprint(writer, 0)
		return
	}

	h := ds.NewMinHeap(n)
	for i := 0; i < n; i++ {
		var x int
		_, _ = fmt.Fscan(reader, &x)
		h.Push(x)
	}

	res := 0
	for h.Size() > 2 {
		a, _ := h.Pop()
		b, _ := h.Pop()
		res += a + b
		h.Push(a + b)
	}

	for h.Size() > 0 {
		a, _ := h.Pop()
		res += a
	}
	_, _ = fmt.Fprint(writer, res)
}
