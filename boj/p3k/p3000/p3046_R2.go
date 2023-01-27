package p3000

import (
	"fmt"
	"io"
)

func Solve3046(reader io.Reader, writer io.Writer) {
	var r1, s int
	_, _ = fmt.Fscan(reader, &r1, &s)

	res := 2*s - r1
	_, _ = fmt.Fprint(writer, res)
}
