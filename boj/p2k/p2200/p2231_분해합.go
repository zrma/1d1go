package p2200

import (
	"fmt"
)

func Solve2231(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	res := findConstructor(n)
	_, _ = fmt.Fprint(writer, res)
}

func findConstructor(n int) int {
	for i := 1; i < n; i++ {
		if decompose(i) == n {
			return i
		}
	}
	return 0
}

func decompose(n int) int {
	res := n
	for ; n > 0; n /= 10 {
		res += n % 10
	}
	return res
}
