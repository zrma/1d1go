package main

import (
	"fmt"

	"1d1go/boj/p8k/p8300"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	res := p8300.Solve8393(n)
	fmt.Println(res)
}
