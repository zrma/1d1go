package p2900

import (
	"fmt"
	"io"
)

func Solve2902(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	_, _ = fmt.Fprint(writer, string(s[0]))
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '-' {
			_, _ = fmt.Fprint(writer, string(s[i+1]))
		}
	}
}
