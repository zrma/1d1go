package p10900

import (
	"fmt"
	"io"
)

func Solve10926(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fprintf(writer, "%s??!", s)
}
