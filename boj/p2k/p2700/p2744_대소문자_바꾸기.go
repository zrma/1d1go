package p2700

import (
	"fmt"
)

func Solve2744(reader Reader, writer Writer) {
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
