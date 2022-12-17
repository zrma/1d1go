package p1000

import (
	"fmt"
)

func Solve1094(reader Reader, writer Writer) {
	var x int
	_, _ = fmt.Fscan(reader, &x)

	res := 0
	for x > 0 {
		res += x % 2
		x /= 2
	}
	_, _ = fmt.Fprint(writer, res)
}

func Solve1094WithBitCount(reader Reader, writer Writer) {
	var x int
	_, _ = fmt.Fscan(reader, &x)

	res := 0
	for x > 0 {
		res += x & 1
		x >>= 1
	}
	_, _ = fmt.Fprint(writer, res)
}
