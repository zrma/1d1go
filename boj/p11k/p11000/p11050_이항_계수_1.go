package p11000

import (
	"fmt"
)

func Solve11050(reader Reader, writer Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	res := binomialCoefficient(n, k)
	_, _ = fmt.Fprint(writer, res)
}

func binomialCoefficient(n, k int) int {
	if n < k {
		return 0
	}
	if n == k {
		return 1
	}
	if k == 0 {
		return 1
	}

	return binomialCoefficient(n-1, k-1) + binomialCoefficient(n-1, k)
}
