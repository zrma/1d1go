package p1000

import (
	"fmt"
	"io"
)

func Solve1085(reader io.Reader, writer io.Writer) {
	var x, y, w, h int
	_, _ = fmt.Fscan(reader, &x, &y, &w, &h)

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
