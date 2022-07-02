package p10900

import (
	"fmt"
)

func Solve10926(reader Reader, writer Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fprintf(writer, "%s??!", s)
}
