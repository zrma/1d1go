package main

import (
	"fmt"

	"1d1go/boj/p1k/p1000"
)

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)

	res := p1000.Solve1000(a, b)
	fmt.Println(res)
}
