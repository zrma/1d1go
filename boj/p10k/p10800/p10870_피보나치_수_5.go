package p10800

import (
	"fmt"
)

func Solve10870(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	res := fibonacci(int64(n))
	_, _ = fmt.Fprint(writer, res)
}

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
