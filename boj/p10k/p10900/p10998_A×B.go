package p10900

import (
	"fmt"
)

func Solve10998(reader Reader, writer Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a*b)
}
