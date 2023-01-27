package p3000

import (
	"fmt"
	"io"
)

func Solve3034(reader io.Reader, writer io.Writer) {
	var n, w, h int
	_, _ = fmt.Fscan(reader, &n, &w, &h)

	limit := w*w + h*h

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		if v*v <= limit {
			_, _ = fmt.Fprintln(writer, "DA")
		} else {
			_, _ = fmt.Fprintln(writer, "NE")
		}
	}
}
