package p2800

import (
	"fmt"
	"io"
	"strings"
)

func Solve2857(reader io.Reader, writer io.Writer) {
	const n = 5

	found := false
	for i := 1; i <= n; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)

		if strings.Contains(s, "FBI") {
			_, _ = fmt.Fprint(writer, i, " ")
			found = true
		}
	}

	if !found {
		_, _ = fmt.Fprint(writer, "HE GOT AWAY!")
	}
}
