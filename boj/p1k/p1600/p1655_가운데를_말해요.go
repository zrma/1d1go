package p1600

import (
	"fmt"
	"io"

	"1d1go/utils/ds"
)

func Solve1655(reader io.Reader, writer io.Writer) {
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

		max, _ := maxH.Peek()
		min, _ := minH.Peek()
		if max > min {
			maxP, _ := maxH.Pop()
			minP, _ := minH.Pop()
			maxH.Push(minP)
			minH.Push(maxP)
		}
		res, _ := maxH.Peek()
		_, _ = fmt.Fprintln(writer, res)
	}
}
