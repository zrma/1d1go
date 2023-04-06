package p9900

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

func Solve9996(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	var s string
	_, _ = fmt.Fscanln(reader, &s)

	r := regexp.MustCompile(`^` + strings.ReplaceAll(s, "*", ".*") + `$`)

	for i := 0; i < n; i++ {
		var t string
		_, _ = fmt.Fscanln(reader, &t)
		if r.MatchString(t) {
			_, _ = fmt.Fprintln(writer, "DA")
		} else {
			_, _ = fmt.Fprintln(writer, "NE")
		}
	}
}
