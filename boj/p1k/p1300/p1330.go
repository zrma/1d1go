package p1300

import (
	"fmt"
	"strconv"
)

func Solve1330(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

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
