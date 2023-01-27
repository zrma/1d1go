package p2700

import (
	"fmt"
	"io"
)

func Solve2742(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := n; i > 0; i-- {
		_, _ = fmt.Fprintln(writer, i)
	}
}
