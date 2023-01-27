package p1000

import (
	"fmt"
	"io"

	"1d1go/utils/integer"
)

func Solve1074(reader io.Reader, writer io.Writer) {
	var n, r, c int
	_, _ = fmt.Fscan(reader, &n, &r, &c)

	length := integer.Pow(2, n)
	_, _ = fmt.Fprint(writer, find(length, r, c))
}

func find(length, r, c int) int {
	if length == 1 {
		return 0
	}

	length /= 2
	if r < length && c < length {
		return find(length, r, c)
	} else if r < length && c >= length {
		return length*length + find(length, r, c-length)
	} else if r >= length && c < length {
		return 2*length*length + find(length, r-length, c)
	} else {
		return 3*length*length + find(length, r-length, c-length)
	}
}
