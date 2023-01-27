package p12900

import (
	"fmt"
	"io"
)

func Solve12904(reader io.Reader, writer io.Writer) {
	var s, t string
	_, _ = fmt.Fscan(reader, &s, &t)

	res := 0
	for len(s) < len(t) {
		if t[len(t)-1] == 'A' {
			t = t[:len(t)-1]
		} else {
			t = t[:len(t)-1]
			t = reverse(t)
		}
		if s == t {
			res = 1
			break
		}
	}
	_, _ = fmt.Fprint(writer, res)
}

func reverse(t string) string {
	res := make([]byte, len(t))
	for i := range res {
		res[i] = t[len(t)-1-i]
	}
	return string(res)
}
