package p10800

import (
	"fmt"
	"io"
)

func Solve10808(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	var cnt [26]int
	for _, c := range s {
		cnt[c-'a']++
	}
	for _, c := range cnt {
		_, _ = fmt.Fprint(writer, c, " ")
	}
}
