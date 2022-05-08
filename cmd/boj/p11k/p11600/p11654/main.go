package main

import (
	"fmt"

	"1d1go/boj/p11k/p11600"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	res := p11600.Solve11654(s)
	fmt.Println(res)
}
