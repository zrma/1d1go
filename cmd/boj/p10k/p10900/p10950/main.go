package main

import (
	"fmt"

	"1d1go/boj/p10k/p10900"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	arr2D := make([][2]int, n)
	for i := range arr2D {
		_, _ = fmt.Scan(&arr2D[i][0], &arr2D[i][1])
	}

	res := p10900.Solve10950(arr2D)
	for _, v := range res {
		fmt.Println(v)
	}
}
