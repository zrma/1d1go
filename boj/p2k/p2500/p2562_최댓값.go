package p2500

import (
	"fmt"
	"io"
	"math"
)

func Solve2562(reader io.Reader, writer io.Writer) {
	const count = 9
	var max = math.MinInt
	var pos int
	for i := 1; i <= count; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		if n > max {
			max = n
			pos = i
		}
	}
	_, _ = fmt.Fprintln(writer, max)
	_, _ = fmt.Fprintln(writer, pos)
}
