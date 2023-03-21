package p27800

import (
	"fmt"
	"io"
)

func Solve27866(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	var i int
	_, _ = fmt.Fscanln(reader, &i)

	_, _ = fmt.Fprint(writer, string(s[i-1]))
}
