package p11900

import (
	"fmt"
	"io"
	"strings"
)

func Solve11944(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscanln(reader, &n, &m)

	buf := new(strings.Builder)
	buf.Grow(m)

	for i := 0; i < n; i++ {
		fmt.Fprintf(buf, "%d", n)
		if buf.Len() >= m {
			break
		}
	}

	res := buf.String()
	if len(res) > m {
		res = res[:m]
	}
	_, _ = fmt.Fprint(writer, res)
}
