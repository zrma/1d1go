package p2700

import (
	"fmt"
)

func Solve2739(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 1; i < 10; i++ {
		_, _ = fmt.Fprintf(writer, "%d * %d = %d\n", n, i, n*i)
	}
}
