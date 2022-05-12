package main

import (
	"fmt"

	"1d1go/boj/p1k/p1100"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	res := p1100.Solve1110(n)
	fmt.Println(res)
}
