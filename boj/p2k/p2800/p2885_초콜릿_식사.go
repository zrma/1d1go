package p2800

import (
	"fmt"
	"io"
)

func Solve2885(reader io.Reader, writer io.Writer) {
	var k int
	_, _ = fmt.Fscan(reader, &k)

	size := 1
	for size < k {
		size *= 2
	}
	fullSize := size

	cnt := 0
	for k > 0 {
		if k >= size {
			k -= size
		} else {
			size /= 2
			cnt++
		}
	}

	_, _ = fmt.Fprint(writer, fullSize, cnt)
}
