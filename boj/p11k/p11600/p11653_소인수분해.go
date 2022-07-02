package p11600

import (
	"fmt"
)

func Solve11653(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	res := factorization(n)
	for _, v := range res {
		_, _ = fmt.Fprintln(writer, v)
	}
}

func factorization(n int) []int {
	if n == 1 {
		return []int{}
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return append(factorization(i), factorization(n/i)...)
		}
	}
	return []int{n}
}
