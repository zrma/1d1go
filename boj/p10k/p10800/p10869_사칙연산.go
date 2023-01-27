package p10800

import (
	"fmt"
	"io"
)

func Solve10869(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)

	_, _ = fmt.Fprintln(writer, a+b)
	_, _ = fmt.Fprintln(writer, a-b)
	_, _ = fmt.Fprintln(writer, a*b)
	_, _ = fmt.Fprintln(writer, a/b)
	_, _ = fmt.Fprintln(writer, a%b)
}
