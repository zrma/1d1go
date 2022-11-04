package p2700

import (
	"fmt"
)

func Solve2752(reader Reader, writer Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	if a > b {
		a, b = b, a
	}

	if b < c {
		_, _ = fmt.Fprint(writer, a, b, c)
	} else {
		if a < c {
			_, _ = fmt.Fprint(writer, a, c, b)
		} else {
			_, _ = fmt.Fprint(writer, c, a, b)
		}
	}
}
