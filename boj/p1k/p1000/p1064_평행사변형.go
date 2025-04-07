package p1000

import (
	"fmt"
	"io"
	"math"
)

func Solve1064(reader io.Reader, writer io.Writer) {
	var ax, ay, bx, by, cx, cy float64
	_, _ = fmt.Fscan(reader, &ax, &ay, &bx, &by, &cx, &cy)

	ab := distance(ax, ay, bx, by)
	bc := distance(bx, by, cx, cy)
	ca := distance(cx, cy, ax, ay)

	if ab == 0 || bc == 0 || ca == 0 {
		_, _ = fmt.Fprint(writer, -1)
		return
	}

	if slope(ax, ay, bx, by) == slope(bx, by, cx, cy) ||
		slope(bx, by, cx, cy) == slope(cx, cy, ax, ay) ||
		slope(cx, cy, ax, ay) == slope(ax, ay, bx, by) {
		_, _ = fmt.Fprint(writer, -1)
		return
	}

	length0 := 2*ab + 2*bc
	length1 := 2*bc + 2*ca
	length2 := 2*ca + 2*ab

	min := math.Min(math.Min(length0, length1), length2)
	max := math.Max(math.Max(length0, length1), length2)

	_, _ = fmt.Fprint(writer, max-min)
}

func distance(x0, y0, x1, y1 float64) float64 {
	return math.Sqrt((x0-x1)*(x0-x1) + (y0-y1)*(y0-y1))
}

func slope(x0, y0, x1, y1 float64) float64 {
	return (y1 - y0) / (x1 - x0)
}
