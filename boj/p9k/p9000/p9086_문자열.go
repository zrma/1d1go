package p9000

import (
	"fmt"
	"io"
)

func Solve9086(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		_, _ = fmt.Fprintf(writer, "%c%c\n", s[0], s[len(s)-1])
	}
}
