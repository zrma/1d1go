package p10700

import (
	"fmt"
	"io"
)

func Solve10798(reader io.Reader, writer io.Writer) {
	var s [5]string
	for i := 0; i < 5; i++ {
		_, _ = fmt.Fscanln(reader, &s[i])
	}

	for i := 0; i < 15; i++ {
		for j := 0; j < 5; j++ {
			if i < len(s[j]) {
				_, _ = fmt.Fprint(writer, string(s[j][i]))
			}
		}
	}
}
