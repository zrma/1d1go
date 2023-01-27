package p3700

import (
	"fmt"
	"io"
)

func Solve3733(reader io.Reader, writer io.Writer) {
	for {
		var n, s int
		_, err := fmt.Fscan(reader, &n, &s)
		if err != nil {
			break
		}

		res := s / (n + 1)
		_, _ = fmt.Fprintln(writer, res)
	}
}
