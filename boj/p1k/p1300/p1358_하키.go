package p1300

import (
	"fmt"
	"math"
	"strconv"
)

func Solve1358(scanner Scanner, writer Writer) {
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	p, _ := strconv.Atoi(scanner.Text())

	cX0 := x
	cY0 := y + h/2

	cX1 := x + w
	cY1 := y + h/2

	r := h / 2

	count := 0
	for i := 0; i < p; i++ {
		scanner.Scan()
		pX, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		pY, _ := strconv.Atoi(scanner.Text())

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
