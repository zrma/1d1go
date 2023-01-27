package p2600

import (
	"fmt"
	"io"
)

func Solve2609(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)

	_, _ = fmt.Fprintln(writer, gcd(a, b))
	_, _ = fmt.Fprintln(writer, lcm(a, b))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
