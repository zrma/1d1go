package p2800

import (
	"fmt"
	"io"
)

func Solve2845(reader io.Reader, writer io.Writer) {
	var l, p int
	_, _ = fmt.Fscan(reader, &l, &p)

	tot := l * p

	var arr [5]int
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	for _, v := range arr {
		_, _ = fmt.Fprintf(writer, "%d ", v-tot)
	}
}
