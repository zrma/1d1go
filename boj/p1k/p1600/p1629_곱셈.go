package p1600

import (
	"fmt"
	"io"
)

func Solve1629(reader io.Reader, writer io.Writer) {
	var a, b, c int64
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	res := pow(a, b, c)
	_, _ = fmt.Fprint(writer, res)
}

func pow(a, b, c int64) int64 {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a % c
	}
	if b%2 == 0 {
		return pow(a*a%c, b/2, c)
	}
	return a * pow(a*a%c, b/2, c) % c
}
