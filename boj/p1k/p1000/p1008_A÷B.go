package p1000

import (
	"fmt"
)

func Solve1008(reader Reader, writer Writer) {
	var a, b float64
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a/b)
}
