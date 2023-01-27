package p4100

import (
	"fmt"
	"io"
)

func Solve4153(reader io.Reader, writer io.Writer) {
	for {
		var a, b, c int
		_, _ = fmt.Fscan(reader, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			return
		}

		if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
			_, _ = fmt.Fprintln(writer, "right")
		} else {
			_, _ = fmt.Fprintln(writer, "wrong")
		}
	}
}
