package p3000

import (
	"fmt"
)

func Solve3046(reader Reader, writer Writer) {
	var r1, s int
	_, _ = fmt.Fscan(reader, &r1, &s)

	res := 2*s - r1
	_, _ = fmt.Fprint(writer, res)
}
