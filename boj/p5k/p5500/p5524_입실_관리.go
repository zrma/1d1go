package p5500

import (
	"bytes"
	"fmt"
	"io"
)

func Solve5524(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	buf := new(bytes.Buffer)

	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)

		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				buf.WriteByte(byte(c) + 32)
			} else {
				buf.WriteByte(byte(c))
			}
		}

		buf.WriteByte('\n')
	}

	_, _ = fmt.Fprint(writer, buf.String())
}
