package p1600

import (
	"fmt"
	"io"
)

func Solve1676(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	res := n/5 + n/25 + n/125
	_, _ = fmt.Fprint(writer, res)
}
