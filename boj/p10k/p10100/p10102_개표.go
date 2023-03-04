package p10100

import (
	"fmt"
	"io"
)

func Solve10102(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var s string
	_, _ = fmt.Fscan(reader, &s)

	a, b := 0, 0
	for _, c := range s {
		if c == 'A' {
			a++
		} else {
			b++
		}
	}

	if a > b {
		_, _ = fmt.Fprint(writer, "A")
	} else if a < b {
		_, _ = fmt.Fprint(writer, "B")
	} else {
		_, _ = fmt.Fprint(writer, "Tie")
	}
}
