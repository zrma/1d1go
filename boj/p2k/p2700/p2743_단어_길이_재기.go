package p2700

import (
	"fmt"
)

func Solve2743(reader Reader, writer Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fprint(writer, len(s))
}
