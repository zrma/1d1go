package p1300

import (
	"fmt"
	"io"
)

func Solve1357(reader io.Reader, writer io.Writer) {
	var x, y int
	_, _ = fmt.Fscan(reader, &x, &y)

	_, _ = fmt.Fprint(writer, reverse(reverse(x)+reverse(y)))
}

func reverse(y int) int {
	var x int
	for y > 0 {
		x = x*10 + y%10
		y /= 10
	}
	return x
}
