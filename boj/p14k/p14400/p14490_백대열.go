package p14400

import (
	"fmt"
	"io"
)

func Solve14490(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscanf(reader, "%d:%d", &n, &m)

	g := gcd(n, m)
	_, _ = fmt.Fprintf(writer, "%d:%d", n/g, m/g)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
