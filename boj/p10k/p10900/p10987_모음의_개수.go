package p10900

import (
	"fmt"
	"io"
)

func Solve10987(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	var cnt int
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			cnt++
		}
	}
	_, _ = fmt.Fprint(writer, cnt)
}
