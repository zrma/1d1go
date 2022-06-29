package p1300

import (
	"fmt"
)

func Solve1330(reader Reader, writer Writer) {
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
