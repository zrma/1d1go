package p2700

import (
	"fmt"
	"io"
)

func Solve2743(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fprint(writer, len(s))
}
