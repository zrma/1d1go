package p10800

import (
	"fmt"
	"io"
)

func Solve10821(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	n := 0
	for i := range s {
		if s[i] == ',' {
			n++
		}
	}
	_, _ = fmt.Fprint(writer, n+1)
}
