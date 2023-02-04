package p1100

import (
	"fmt"
	"io"
)

func Solve1100(reader io.Reader, writer io.Writer) {
	const size = 8

	res := 0
	for i := 0; i < size; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		for j, c := range s {
			if c == '.' {
				continue
			}
			res += (i + j + 1) % 2
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
