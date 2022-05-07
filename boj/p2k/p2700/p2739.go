package p2700

import (
	"fmt"
)

func Solve2739(n int) {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}
