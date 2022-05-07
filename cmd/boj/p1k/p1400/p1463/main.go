package main

import (
	"fmt"

	"1d1go/boj/p1k/p1400"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	res := p1400.Solve1463(n)
	fmt.Println(res)
}
