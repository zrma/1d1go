package p2800

import (
	"fmt"
)

func Solve2869(reader Reader, writer Writer) {
	var a, b, v int
	_, _ = fmt.Fscan(reader, &a, &b, &v)

	if v <= a {
		_, _ = fmt.Fprint(writer, 1)
		return
	}

	res := (v-b-1)/(a-b) + 1
	_, _ = fmt.Fprint(writer, res)
}
