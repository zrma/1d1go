package p2700

import (
	"fmt"
)

func Solve2742(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := n; i > 0; i-- {
		_, _ = fmt.Fprintln(writer, i)
	}
}
