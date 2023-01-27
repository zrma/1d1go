package p3000

import (
	"fmt"
	"io"
)

func Solve3036(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	var arr0 int
	_, _ = fmt.Fscan(reader, &arr0)

	for i := 1; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		div := gcd(arr0, v)
		_, _ = fmt.Fprintf(writer, "%d/%d\n", arr0/div, v/div)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
