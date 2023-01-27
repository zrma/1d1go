package p2400

import (
	"fmt"
	"io"
)

func Solve2480(reader io.Reader, writer io.Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	price := calcPrice(a, b, c)

	_, _ = fmt.Fprint(writer, price)
}

func calcPrice(a, b, c int) int {
	if a == b && b == c {
		return 10_000 + a*1_000
	}
	if a == b || a == c {
		return 1_000 + a*100
	}
	if b == c {
		return 1_000 + b*100
	}
	var max = a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max * 100
}
