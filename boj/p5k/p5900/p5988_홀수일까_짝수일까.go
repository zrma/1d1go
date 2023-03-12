package p5900

import (
	"fmt"
	"io"
)

func Solve5988(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)
		if s[len(s)-1]%2 == 0 {
			_, _ = fmt.Fprintln(writer, "even")
		} else {
			_, _ = fmt.Fprintln(writer, "odd")
		}
	}
}
