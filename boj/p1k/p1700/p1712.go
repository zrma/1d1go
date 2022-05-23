package p1700

import (
	"fmt"
	"strconv"
)

func Solve1712(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	res := calcBreakEvenPoint(a, b, c)
	_, _ = fmt.Fprint(writer, res)
}

func calcBreakEvenPoint(a, b, c int) int {
	if b >= c {
		return -1
	}
	return a/(c-b) + 1
}
