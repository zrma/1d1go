package p1200

import (
	"fmt"
	"io"
)

func Solve1225(reader io.Reader, writer io.Writer) {
	var a, b string
	_, _ = fmt.Fscan(reader, &a, &b)

	var sum int
	for _, c0 := range a {
		for _, c1 := range b {
			sum += int(c0-'0') * int(c1-'0')
		}
	}
	_, _ = fmt.Fprint(writer, sum)
}
