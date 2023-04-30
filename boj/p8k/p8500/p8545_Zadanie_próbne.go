package p8500

import (
	"fmt"
	"io"
)

func Solve8545(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	for i := 0; i < len(s); i++ {
		_, _ = fmt.Fprintf(writer, "%c", s[len(s)-i-1])
	}
}
