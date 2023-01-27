package p1000

import (
	"fmt"
	"io"
)

func Solve1000(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a+b)
}
