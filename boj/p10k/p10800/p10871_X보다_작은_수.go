package p10800

import (
	"fmt"
)

func Solve10871(reader Reader, writer Writer) {
	var n, x int
	_, _ = fmt.Fscan(reader, &n, &x)

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		if v < x {
			_, _ = fmt.Fprintf(writer, "%d ", v)
		}
	}
}
