package p1200

import (
	"fmt"
	"io"
	"strings"
)

func Solve1251(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	tmp := new(strings.Builder)
	tmp.Grow(len(s))

	var res string
	for i := 0; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			tmp.Reset()
			for k := i; k >= 0; k-- {
				_, _ = fmt.Fprint(tmp, string(s[k]))
			}
			for k := j; k > i; k-- {
				_, _ = fmt.Fprint(tmp, string(s[k]))
			}
			for k := len(s) - 1; k > j; k-- {
				_, _ = fmt.Fprint(tmp, string(s[k]))
			}
			if res == "" || res > tmp.String() {
				res = tmp.String()
			}
		}
	}

	_, _ = fmt.Fprint(writer, res)
}
