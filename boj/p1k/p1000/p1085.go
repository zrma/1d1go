package p1000

import (
	"fmt"
	"strconv"
)

func Solve1085(scanner Scanner, writer Writer) {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	h, _ := strconv.Atoi(scanner.Text())

	distance := x
	if distance > y {
		distance = y
	}
	if distance > w-x {
		distance = w - x
	}
	if distance > h-y {
		distance = h - y
	}

	_, _ = fmt.Fprint(writer, distance)
}
