package main

import (
	"fmt"

	"1d1go/boj/p2k/p2400"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		var a int
		_, _ = fmt.Scan(&a)
		arr = append(arr, a)
	}
	res := p2400.Solve2437(arr)
	fmt.Println(res)
}
