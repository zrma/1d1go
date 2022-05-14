package p1000

import (
	"fmt"
	"strconv"
)

func Solve1008(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	b, _ := strconv.ParseFloat(scanner.Text(), 64)

	_, _ = fmt.Fprintln(writer, a/b)
}
