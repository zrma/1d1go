package p2700

import (
	"fmt"
	"strconv"
)

func Solve2739(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 1; i < 10; i++ {
		_, _ = fmt.Fprintf(writer, "%d * %d = %d\n", n, i, n*i)
	}
}
