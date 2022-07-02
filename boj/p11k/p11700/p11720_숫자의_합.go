package p11700

import (
	"fmt"
)

func Solve11720(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var s string
	_, _ = fmt.Fscan(reader, &s)

	var sum int
	for i := 0; i < n && i < len(s); i++ {
		c := int(s[i] - '0')
		sum += c
	}
	_, _ = fmt.Fprintf(writer, "%d", sum)
}
