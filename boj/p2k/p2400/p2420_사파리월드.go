package p2400

import (
	"fmt"
	"io"
)

func Solve2420(reader io.Reader, writer io.Writer) {
	var n, m int64
	_, _ = fmt.Fscan(reader, &n, &m)

	diff := n - m
	if diff < 0 {
		diff = -diff
	}
	_, _ = fmt.Fprint(writer, diff)
}
