package p2700

import (
	"fmt"
	"io"
)

func Solve2741(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fprintln(writer, i)
	}
}
