package p1000

import (
	"fmt"
	"io"
	"regexp"
)

func Solve1013(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscanln(reader, &t)

	reg := regexp.MustCompile(`^(100+1+|01)+$`)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)
		if reg.MatchString(s) {
			_, _ = fmt.Fprintln(writer, "YES")
		} else {
			_, _ = fmt.Fprintln(writer, "NO")
		}
	}
}
