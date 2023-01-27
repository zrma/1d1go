package p2700

import (
	"fmt"
	"io"
)

func Solve2753(reader io.Reader, writer io.Writer) {
	var year int
	_, _ = fmt.Fscan(reader, &year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		_, _ = fmt.Fprint(writer, "1")
	} else {
		_, _ = fmt.Fprint(writer, "0")
	}
}
