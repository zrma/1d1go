package p10900

import (
	"fmt"
)

func Solve10951(reader Reader, writer Writer) {
	for {
		var a, b int
		if _, err := fmt.Fscan(reader, &a, &b); err != nil {
			break
		}
		_, _ = fmt.Fprintln(writer, a+b)
	}
}
