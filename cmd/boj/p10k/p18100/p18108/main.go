package main

import (
	"fmt"

	"1d1go/boj/p18k/p18108"
)

func main() {
	var year int
	_, _ = fmt.Scan(&year)

	res := p18108.Solve18108(year)
	fmt.Println(res)
}
