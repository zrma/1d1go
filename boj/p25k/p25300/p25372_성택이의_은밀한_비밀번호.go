package p25300

import (
	"fmt"
	"io"
)

func Solve25372(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		if len(s) >= 6 && len(s) <= 9 {
			_, _ = fmt.Fprintln(writer, "yes")
		} else {
			_, _ = fmt.Fprintln(writer, "no")
		}
	}
}
