package p4200

import (
	"fmt"
	"io"
)

func Solve4299(reader io.Reader, writer io.Writer) {
	var sum, diff int
	_, _ = fmt.Fscan(reader, &sum, &diff)

	if sum < diff {
		_, _ = fmt.Fprint(writer, -1)
		return
	}

	a := sum + (diff-sum)/2
	b := (sum - diff) / 2

	if a+b != sum || a-b != diff {
		_, _ = fmt.Fprint(writer, -1)
		return
	}
	_, _ = fmt.Fprintf(writer, "%d %d", a, b)
}
