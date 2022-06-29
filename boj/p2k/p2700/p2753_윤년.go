package p2700

import (
	"fmt"
)

func Solve2753(reader Reader, writer Writer) {
	var year int
	_, _ = fmt.Fscan(reader, &year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		_, _ = fmt.Fprint(writer, "1")
	} else {
		_, _ = fmt.Fprint(writer, "0")
	}
}
