package p1100

import (
	"fmt"
	"io"
)

func Solve1193(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	sum := 0
	i := 1
	for ; sum < n; i++ {
		sum += i
	}

	if i%2 == 0 {
		j := sum - n
		_, _ = fmt.Fprintf(writer, "%d/%d", j+1, i-j-1)
	} else {
		j := sum - n
		_, _ = fmt.Fprintf(writer, "%d/%d", i-j-1, j+1)
	}
}
