package p11500

import (
	"fmt"
	"io"
)

func Solve11586(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscanln(reader, &n)

	var mirror []string
	for i := 0; i < n; i++ {
		var line string
		_, _ = fmt.Fscanln(reader, &line)
		mirror = append(mirror, line)
	}

	var k int
	_, _ = fmt.Fscanln(reader, &k)

	switch k {
	case 2:
		for i := 0; i < n; i++ {
			for j := n - 1; j >= 0; j-- {
				_, _ = fmt.Fprintf(writer, "%c", mirror[i][j])
			}
			_, _ = fmt.Fprintln(writer)
		}
	case 3:
		for i := n - 1; i >= 0; i-- {
			_, _ = fmt.Fprintln(writer, mirror[i])
		}
	default:
		for _, line := range mirror {
			_, _ = fmt.Fprintln(writer, line)
		}
	}
}
