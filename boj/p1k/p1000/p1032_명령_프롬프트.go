package p1000

import (
	"fmt"
	"io"
)

func Solve1032(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var s string
	_, _ = fmt.Fscan(reader, &s)

	buf := make([]byte, len(s))
	copy(buf, s)

	for i := 1; i < n; i++ {
		_, _ = fmt.Fscan(reader, &s)
		for j := range buf {
			if buf[j] != s[j] {
				buf[j] = '?'
			}
		}
	}

	_, _ = fmt.Fprint(writer, string(buf))
}
