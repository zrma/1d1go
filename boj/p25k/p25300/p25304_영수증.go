package p25300

import (
	"fmt"
	"io"
)

func Solve25304(reader io.Reader, writer io.Writer) {
	var x, n int
	_, _ = fmt.Fscan(reader, &x, &n)

	var sum int
	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		sum += a * b
	}

	if sum == x {
		_, _ = fmt.Fprint(writer, "Yes")
	} else {
		_, _ = fmt.Fprint(writer, "No")
	}
}
