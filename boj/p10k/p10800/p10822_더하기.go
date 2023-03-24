package p10800

import (
	"fmt"
	"io"
	"strings"
)

func Solve10822(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	ss := strings.Split(s, ",")

	var res int
	for _, s := range ss {
		var n int
		_, _ = fmt.Sscan(s, &n)
		res += n
	}

	_, _ = fmt.Fprint(writer, res)
}
