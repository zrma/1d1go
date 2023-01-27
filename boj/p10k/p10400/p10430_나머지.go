package p10400

import (
	"fmt"
	"io"
)

func Solve10430(reader io.Reader, writer io.Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	_, _ = fmt.Fprintln(writer, (a+b)%c)
	_, _ = fmt.Fprintln(writer, (a%c+b%c)%c)
	_, _ = fmt.Fprintln(writer, (a*b)%c)
	_, _ = fmt.Fprintln(writer, (a%c*b%c)%c)
}
