package p1000

import (
	"fmt"
	"math"
)

func Solve1004(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		countOrbitInOut(reader, writer)
	}
}

func countOrbitInOut(reader Reader, writer Writer) {
	var x0, y0, x1, y1, n, x, y, r int

	_, _ = fmt.Fscan(reader, &x0, &y0, &x1, &y1, &n)

	count := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &x, &y, &r)

		contained0 := isInCircle(x0, y0, x, y, r)
		contained1 := isInCircle(x1, y1, x, y, r)

		if contained0 && !contained1 || !contained0 && contained1 {
			count++
		}
	}
	_, _ = fmt.Fprintln(writer, count)
}

func isInCircle(x, y, cX, cY, r int) bool {
	return math.Pow(float64(x-cX), 2)+math.Pow(float64(y-cY), 2) <= math.Pow(float64(r), 2)
}
