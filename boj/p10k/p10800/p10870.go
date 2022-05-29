package p10800

import (
	"fmt"
	"strconv"
)

func Solve10870(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	res := fibonacci(int64(n))
	_, _ = fmt.Fprint(writer, res)
}

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
