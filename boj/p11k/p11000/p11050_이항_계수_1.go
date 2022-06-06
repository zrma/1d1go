package p11000

import (
	"fmt"
	"strconv"
)

func Solve11050(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

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
