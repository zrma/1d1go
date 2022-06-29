package p10800

import (
	"fmt"
)

func Solve10869(reader Reader, writer Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)

	_, _ = fmt.Fprintln(writer, a+b)
	_, _ = fmt.Fprintln(writer, a-b)
	_, _ = fmt.Fprintln(writer, a*b)
	_, _ = fmt.Fprintln(writer, a/b)
	_, _ = fmt.Fprintln(writer, a%b)
}
