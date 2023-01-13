package p2500

import (
	"fmt"
)

func Solve2558(reader Reader, writer Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a+b)
}
