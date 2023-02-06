package p11600

import (
	"fmt"
	"io"
	"sort"
)

func Solve11656(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	n := len(s)
	prefixes := make([]string, n)
	for i := 0; i < n; i++ {
		prefixes[i] = s[i:]
	}

	sort.Strings(prefixes)
	for _, prefix := range prefixes {
		_, _ = fmt.Fprintln(writer, prefix)
	}
}
