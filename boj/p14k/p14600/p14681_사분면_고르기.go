package p14600

import (
	"fmt"
	"io"
)

func Solve14681(reader io.Reader, writer io.Writer) {
	var x, y int
	_, _ = fmt.Fscan(reader, &x, &y)

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
