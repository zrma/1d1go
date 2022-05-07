package main

import (
	"fmt"

	"1d1go/boj/p14k/p14600"
)

func main() {
	var x, y int
	_, _ = fmt.Scan(&x, &y)

	got := p14600.Solve14681(x, y)
	fmt.Println(got)
}
