package main

import (
	"fmt"

	"1d1go/boj/p2k/p2500"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	res := p2500.Solve2588(a, b)
	for _, n := range res {
		fmt.Println(n)
	}
}
