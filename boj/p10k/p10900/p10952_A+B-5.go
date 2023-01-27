package p10900

import (
	"fmt"
	"io"
)

func Solve10952(reader io.Reader, writer io.Writer) {
	for {
		var a, b int
		if _, err := fmt.Fscan(reader, &a, &b); err != nil {
			break
		}
		if a == 0 && b == 0 {
			break
		}

		_, _ = fmt.Fprintln(writer, a+b)
	}
}
