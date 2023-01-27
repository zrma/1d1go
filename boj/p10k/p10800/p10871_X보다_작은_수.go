package p10800

import (
	"fmt"
	"io"
)

func Solve10871(reader io.Reader, writer io.Writer) {
	var n, x int
	_, _ = fmt.Fscan(reader, &n, &x)

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		if v < x {
			_, _ = fmt.Fprintf(writer, "%d ", v)
		}
	}
}
