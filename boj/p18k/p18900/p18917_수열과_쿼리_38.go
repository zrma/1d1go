package p18900

import (
	"fmt"
	"io"
)

func Solve18917(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var sum int64
	var xor int64
	for i := 0; i < n; i++ {
		var q int64
		_, _ = fmt.Fscan(reader, &q)

		switch q {
		case 1:
			var v int64
			_, _ = fmt.Fscan(reader, &v)
			sum += v
			xor ^= v
		case 2:
			var v int64
			_, _ = fmt.Fscan(reader, &v)
			sum -= v
			xor ^= v
		case 3:
			_, _ = fmt.Fprintln(writer, sum)
		default:
			_, _ = fmt.Fprintln(writer, xor)
		}
	}
}
