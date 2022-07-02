package p8300

import (
	"fmt"
)

func Solve8393(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	_, _ = fmt.Fprint(writer, sum)
}

func Solve8393AP(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	sum := n * (n + 1) / 2
	_, _ = fmt.Fprint(writer, sum)
}
