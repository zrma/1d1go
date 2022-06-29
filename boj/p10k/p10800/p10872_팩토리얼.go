package p10800

import (
	"fmt"
)

func Solve10872(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	got := factorial(int64(n))
	_, _ = fmt.Fprint(writer, got)
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
