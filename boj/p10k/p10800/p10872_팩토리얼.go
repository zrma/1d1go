package p10800

import (
	"fmt"
	"strconv"
)

func Solve10872(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	got := factorial(int64(n))
	_, _ = fmt.Fprint(writer, got)
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
