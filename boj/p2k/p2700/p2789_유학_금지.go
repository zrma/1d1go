package p2700

import (
	"fmt"
	"io"
)

func Solve2789(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	for _, c := range s {
		switch c {
		case 'C', 'A', 'M', 'B', 'R', 'I', 'D', 'G', 'E':
			continue
		}
		_, _ = fmt.Fprintf(writer, "%c", c)
	}
}
