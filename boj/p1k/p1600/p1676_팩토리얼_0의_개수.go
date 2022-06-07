package p1600

import (
	"fmt"
	"strconv"
)

func Solve1676(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	res := n/5 + n/25 + n/125
	_, _ = fmt.Fprint(writer, res)
}
