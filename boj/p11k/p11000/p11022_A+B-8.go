package p11000

import (
	"fmt"
	"io"
)

func Solve11022(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		_, _ = fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}
}
