package p1400

import (
	"fmt"
	"io"
	"sort"
)

func Solve1427(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	_, _ = fmt.Fprint(writer, string(b))
}
