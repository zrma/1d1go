package p10900

import (
	"fmt"
	"io"
)

func Solve10998(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a*b)
}
