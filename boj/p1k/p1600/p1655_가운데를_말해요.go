package p1600

import (
	"fmt"

	"1d1go/utils/ds"
)

func Solve1655(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	maxH := ds.NewMaxHeap(n)
	minH := ds.NewMinHeap(n)

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if maxH.Size() == minH.Size() {
			maxH.Push(v)
		} else {
			minH.Push(v)
		}

		if maxH.Peek() > minH.Peek() {
			maxP, minP := maxH.Pop(), minH.Pop()
			maxH.Push(minP)
			minH.Push(maxP)
		}
		_, _ = fmt.Fprintln(writer, maxH.Peek())
	}
}
