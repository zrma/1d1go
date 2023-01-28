package p11700

import (
	"fmt"
	"io"
)

func Solve11721(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	for i := 0; i < len(s); i += 10 {
		if i+10 < len(s) {
			_, _ = fmt.Fprintln(writer, s[i:i+10])
		} else {
			_, _ = fmt.Fprintln(writer, s[i:])
		}
	}
}
