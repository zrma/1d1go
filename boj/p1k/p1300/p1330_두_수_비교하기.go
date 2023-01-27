package p1300

import (
	"fmt"
	"io"
)

func Solve1330(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, comp(a, b))
}

func comp(a, b int) string {
	if a > b {
		return ">"
	}
	if a < b {
		return "<"
	}
	return "=="
}
