package p10400

import (
	"fmt"
)

func Solve10430(reader Reader, writer Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	_, _ = fmt.Fprintln(writer, (a+b)%c)
	_, _ = fmt.Fprintln(writer, (a%c+b%c)%c)
	_, _ = fmt.Fprintln(writer, (a*b)%c)
	_, _ = fmt.Fprintln(writer, (a%c*b%c)%c)
}
