package p4600

import (
	"fmt"
)

func Solve4673(writer Writer) {
	const n = 10000

	arr := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if arr[i] {
			continue
		}
		for j := i; j <= n; {
			if j = decompose(j); j <= n {
				arr[j] = true
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !arr[i] {
			_, _ = fmt.Fprintln(writer, i)
		}
	}
}

func decompose(n int) int {
	res := n
	for ; n > 0; n /= 10 {
		res += n % 10
	}
	return res
}
