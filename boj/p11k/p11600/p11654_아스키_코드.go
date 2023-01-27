package p11600

import (
	"fmt"
	"io"
)

func Solve11654(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fprint(writer, int(s[0]))
}
