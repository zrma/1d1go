package p1700

import (
	"fmt"
	"io"
)

func Solve1712(reader io.Reader, writer io.Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	res := calcBreakEvenPoint(a, b, c)
	_, _ = fmt.Fprint(writer, res)
}

func calcBreakEvenPoint(a, b, c int) int {
	if b >= c {
		return -1
	}
	return a/(c-b) + 1
}
