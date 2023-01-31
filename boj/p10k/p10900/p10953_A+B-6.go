package p10900

import (
	"fmt"
	"io"
)

func Solve10953(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		var a, b int
		_, _ = fmt.Sscanf(s, "%d,%d", &a, &b)
		_, _ = fmt.Fprintln(writer, a+b)
	}
}
