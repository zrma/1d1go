package p2400

import (
	"fmt"
)

func Solve2475(reader Reader, writer Writer) {
	const n = 5

	sum := 0
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		sum += v * v
	}
	res := sum % 10
	_, _ = fmt.Fprint(writer, res)
}
