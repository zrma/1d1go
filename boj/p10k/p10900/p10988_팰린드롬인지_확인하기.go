package p10900

import (
	"fmt"
	"io"
)

func Solve10988(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			_, _ = fmt.Fprint(writer, 0)
			return
		}
	}
	_, _ = fmt.Fprint(writer, 1)
}
