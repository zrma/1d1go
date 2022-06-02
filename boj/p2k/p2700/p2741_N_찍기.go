package p2700

import (
	"fmt"
	"strconv"
)

func Solve2741(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= n; i++ {
		_, _ = fmt.Fprintln(writer, i)
	}
}
