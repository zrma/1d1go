package p1300

import (
	"fmt"
	"math"
)

func Solve1358(reader Reader, writer Writer) {
	var w, h, x, y, p int
	_, _ = fmt.Fscan(reader, &w, &h, &x, &y, &p)

	cX0 := x
	cY0 := y + h/2

	cX1 := x + w
	cY1 := y + h/2

	r := h / 2

	count := 0
	for i := 0; i < p; i++ {
		var pX, pY int
		_, _ = fmt.Fscan(reader, &pX, &pY)

		if isInCircle(pX, pY, cX0, cY0, r) || isInCircle(pX, pY, cX1, cY1, r) {
			count++
		} else if isInRect(pX, pY, x, y, x+w, y+h) {
			count++
		}
	}
	_, _ = fmt.Fprint(writer, count)
}

func isInCircle(x, y, cX, cY, r int) bool {
	return math.Pow(float64(x-cX), 2)+math.Pow(float64(y-cY), 2) <= math.Pow(float64(r), 2)
}

func isInRect(x, y, rX0, rY0, rX1, rY1 int) bool {
	return x >= rX0 && x <= rX1 && y >= rY0 && y <= rY1
}
