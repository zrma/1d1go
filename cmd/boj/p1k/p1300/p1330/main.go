package main

import (
	"fmt"

	"1d1go/boj/p1k/p1300"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	res := p1300.Solve1330(a, b)
	fmt.Println(res)
}
