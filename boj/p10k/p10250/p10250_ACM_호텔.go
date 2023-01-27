package p10250

import (
	"fmt"
	"io"
)

func Solve10250(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var h, w, k int
		_, _ = fmt.Fscan(reader, &h, &w, &k)

		res := calcNearestRoom(h, w, k)
		_, _ = fmt.Fprintln(writer, res)
	}
}

func calcNearestRoom(h, _, n int) int {
	if n == 1 {
		return 101
	}
	col := n/h + 1
	row := n % h
	if row == 0 {
		row = h
		col -= 1
	}
	return row*100 + col
}
