package main

import (
	"fmt"

	"1d1go/boj/p10k/p10900"
)

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)

	res := p10900.Solve10998(a, b)
	fmt.Println(res)
}
