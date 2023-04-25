package p3000

import (
	"fmt"
	"io"
)

func Solve3028(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscanln(reader, &s)

	var cup = [3]bool{true, false, false}
	for _, c := range s {
		switch c {
		case 'A':
			cup[0], cup[1] = cup[1], cup[0]
		case 'B':
			cup[1], cup[2] = cup[2], cup[1]
		case 'C':
			cup[0], cup[2] = cup[2], cup[0]
		}
	}

	for i, b := range cup {
		if b {
			_, _ = fmt.Fprint(writer, i+1)
			return
		}
	}
}
