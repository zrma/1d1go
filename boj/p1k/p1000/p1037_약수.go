package p1000

import (
	"fmt"
	"math"
)

func Solve1037(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	_, _ = fmt.Fprint(writer, min*max)
}
