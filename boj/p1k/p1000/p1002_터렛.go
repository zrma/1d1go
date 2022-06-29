package p1000

import (
	"fmt"
	"math"
)

func Solve1002(reader Reader, writer Writer) {
	var n, x0, y0, r0, x1, y1, r1 int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &x0, &y0, &r0, &x1, &y1, &r1)

		res := countJunctionOfCircles(x0, y0, r0, x1, y1, r1)
		_, _ = fmt.Fprintln(writer, res)
	}
}

func countJunctionOfCircles(x0, y0, r0, x1, y1, r1 int) int {
	// 동심
	if x0 == x1 && y0 == y1 {
		if r0 == r1 {
			// 동일한 두 원 - 무한히 겹침
			return -1
		} else {
			// 동심원
			return 0
		}
	}
	fX0 := float64(x0)
	fY0 := float64(y0)
	fX1 := float64(x1)
	fY1 := float64(y1)
	fR0 := float64(r0)
	fR1 := float64(r1)

	distance := math.Sqrt(math.Pow(fX1-fX0, 2) + math.Pow(fY1-fY0, 2))
	// 포함
	if distance < math.Abs(fR0-fR1) {
		return 0
	}
	// 외부
	if distance > fR0+fR1 {
		return 0
	}
	// 내접
	if distance == math.Abs(fR0-fR1) {
		return 1
	}
	// 외접
	if distance == fR0+fR1 {
		return 1
	}
	return 2
}
