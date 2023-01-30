package p4900

import (
	"fmt"
	"io"
)

func Solve4999(reader io.Reader, writer io.Writer) {
	var a, b string
	_, _ = fmt.Fscan(reader, &a, &b)

	if len(a) >= len(b) {
		_, _ = fmt.Fprint(writer, "go")
	} else {
		_, _ = fmt.Fprint(writer, "no")
	}
}
