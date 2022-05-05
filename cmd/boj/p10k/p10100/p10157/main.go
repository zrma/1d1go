package main

import (
	"fmt"

	"1d1go/boj/p10k/p10100"
)

func main() {
	var c, r, k int
	_, _ = fmt.Scanf("%d %d", &c, &r)
	_, _ = fmt.Scanf("%d", &k)

	x, y, ok := p10100.Solve10157(c, r, k)
	if ok {
		fmt.Printf("%d %d\n", x, y)
	} else {
		fmt.Println("0")
	}
}
