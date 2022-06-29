package p1000

import (
	"fmt"
)

func Solve1000(reader Reader, writer Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	_, _ = fmt.Fprint(writer, a+b)
}
