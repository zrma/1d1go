package p10900

import (
	"fmt"
	"io"
)

func Solve10950(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		_, _ = fmt.Fprintln(writer, a+b)
	}
}
