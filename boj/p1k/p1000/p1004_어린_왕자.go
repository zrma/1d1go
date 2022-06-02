package p1000

import (
	"fmt"
	"math"
	"strconv"
)

func Solve1004(scanner Scanner, writer Writer) {
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		countOrbitInOut(scanner, writer)
	}
}

func countOrbitInOut(scanner Scanner, writer Writer) {
	scanner.Scan()
	x0, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y0, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	x1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	count := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())

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
