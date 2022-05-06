package main

import (
	"fmt"

	"1d1go/boj/p10k/p10900"
)

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)

	res := p10900.Solve10926(s)
	fmt.Println(res)
}
