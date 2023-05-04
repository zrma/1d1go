package p5500

import (
	"fmt"
	"io"
	"strings"
)

func Solve5555(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	var n int
	_, _ = fmt.Fscanln(reader, &n)

	count := 0
	for i := 0; i < n; i++ {
		var t string
		_, _ = fmt.Fscanln(reader, &t)

		t = t + t
		if strings.Contains(t, s) {
			count++
		}
	}

	_, _ = fmt.Fprint(writer, count)
}
