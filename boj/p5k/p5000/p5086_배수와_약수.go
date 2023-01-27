package p5000

import (
	"fmt"
	"io"
)

func Solve5086(reader io.Reader, writer io.Writer) {
	for {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)

		if a == 0 && b == 0 {
			break
		}

		if b%a == 0 {
			_, _ = fmt.Fprintln(writer, "factor")
		} else if a%b == 0 {
			_, _ = fmt.Fprintln(writer, "multiple")
		} else {
			_, _ = fmt.Fprintln(writer, "neither")
		}
	}
}
