package main

import (
	"fmt"

	"1d1go/boj/p2k/p2700"
)

func main() {
	var year int
	_, _ = fmt.Scan(&year)

	res := p2700.Solve2753(year)
	fmt.Println(res)
}
