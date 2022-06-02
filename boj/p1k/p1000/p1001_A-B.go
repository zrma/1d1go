package p1000

import (
	"fmt"
	"strconv"
)

func Solve1001(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	_, _ = fmt.Fprint(writer, a-b)
}
