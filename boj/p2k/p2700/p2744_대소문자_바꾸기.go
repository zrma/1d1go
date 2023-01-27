package p2700

import (
	"fmt"
	"io"
)

func Solve2744(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			_, _ = fmt.Fprintf(writer, "%c", v-'a'+'A')
		} else {
			_, _ = fmt.Fprintf(writer, "%c", v-'A'+'a')
		}
	}
}
