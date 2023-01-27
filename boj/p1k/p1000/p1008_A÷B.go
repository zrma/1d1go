package p1000

import (
	"fmt"
	"io"
)

func Solve1008(reader io.Reader, writer io.Writer) {
	var a, b float64
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a/b)
}
