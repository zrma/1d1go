package p16900

import (
	"fmt"
	"io"
	"strings"
)

func Solve16916(reader io.Reader, writer io.Writer) {
	var s, p string
	_, _ = fmt.Fscanln(reader, &s)
	_, _ = fmt.Fscanln(reader, &p)

	if strings.Contains(s, p) {
		_, _ = fmt.Fprint(writer, 1)
	} else {
		_, _ = fmt.Fprint(writer, 0)
	}
}
