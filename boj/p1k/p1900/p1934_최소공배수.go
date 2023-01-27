package p1900

import (
	"fmt"
	"io"
)

func Solve1934(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		res := lcm(a, b)
		_, _ = fmt.Fprintln(writer, res)
	}
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
