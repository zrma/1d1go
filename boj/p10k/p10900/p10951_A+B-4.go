package p10900

import (
	"fmt"
	"io"
)

func Solve10951(reader io.Reader, writer io.Writer) {
	for {
		var a, b int
		if _, err := fmt.Fscan(reader, &a, &b); err != nil {
			break
		}
		_, _ = fmt.Fprintln(writer, a+b)
	}
}
