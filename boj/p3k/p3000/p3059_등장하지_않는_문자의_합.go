package p3000

import (
	"fmt"
	"io"
)

func Solve3059(reader io.Reader, writer io.Writer) {
	var t int
	_, _ = fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Fscanln(reader, &s)

		var arr [26]bool

		for _, c := range s {
			arr[c-'A'] = true
		}

		sum := 0
		for i, b := range arr {
			if !b {
				sum += i + 'A'
			}
		}

		_, _ = fmt.Fprintln(writer, sum)
	}
}
