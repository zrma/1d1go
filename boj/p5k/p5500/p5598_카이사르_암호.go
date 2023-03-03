package p5500

import (
	"fmt"
	"io"
)

func Solve5598(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	for _, c := range s {
		_, _ = fmt.Fprintf(writer, "%c", (c-'A'-3+26)%26+'A')
	}
}
