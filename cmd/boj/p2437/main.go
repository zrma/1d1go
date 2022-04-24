package main

import (
	"fmt"

	"1d1go/boj"
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
	res := boj.Solve2437(arr)
	fmt.Println(res)
}
