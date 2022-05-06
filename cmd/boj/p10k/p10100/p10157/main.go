package main

import (
	"fmt"

	"1d1go/boj/p10k/p10100"
)

func main() {
	var c, r, k int
	_, _ = fmt.Scan(&c, &r)
	_, _ = fmt.Scan(&k)

	x, y, ok := p10100.Solve10157(c, r, k)
	if ok {
		fmt.Printf("%d %d\n", x, y)
	} else {
		fmt.Println("0")
	}
}
