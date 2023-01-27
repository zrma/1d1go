package p11400

import (
	"fmt"
	"io"
)

func Solve11478(reader io.Reader, writer io.Writer) {
	m := make(map[string]bool)

	var s string
	_, _ = fmt.Fscan(reader, &s)

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			m[s[i:j]] = true
		}
	}
	_, _ = fmt.Fprint(writer, len(m))
}
