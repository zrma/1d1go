package p2500

import (
	"fmt"
)

func Solve2566(reader Reader, writer Writer) {
	const n = 9

	var maxVal int
	var maxRow, maxCol int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var v int
			_, _ = fmt.Fscan(reader, &v)
			if v > maxVal {
				maxVal = v
				maxRow = i
				maxCol = j
			}
		}
	}

	_, _ = fmt.Fprintln(writer, maxVal)
	_, _ = fmt.Fprintf(writer, "%d %d", maxRow+1, maxCol+1)
}
