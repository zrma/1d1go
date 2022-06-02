package p14600

import (
	"fmt"
	"strconv"
)

func Solve14681(scanner Scanner, writer Writer) {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ := strconv.Atoi(scanner.Text())

	res := posToQuadrant(x, y)
	_, _ = fmt.Fprint(writer, res)
}

func posToQuadrant(x, y int) int {
	if x == 0 || y == 0 {
		return -1
	}
	if x > 0 {
		if y > 0 {
			return 1
		} else {
			return 4
		}
	} else {
		if y > 0 {
			return 2
		} else {
			return 3
		}
	}
}
