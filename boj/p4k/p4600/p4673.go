package p4600

import (
	"fmt"
)

func Solve4673(n int) {
	arr := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if arr[i] {
			continue
		}
		for j := i; j <= n; {
			if j = d(j); j <= n {
				arr[j] = true
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !arr[i] {
			fmt.Println(i)
		}
	}
}

func d(n int) int {
	res := n
	for ; n > 0; n /= 10 {
		res += n % 10
	}
	return res
}
