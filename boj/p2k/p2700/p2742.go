package p2700

import (
	"fmt"
	"strconv"
)

func Solve2742(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := n; i > 0; i-- {
		_, _ = fmt.Fprintln(writer, i)
	}
}
