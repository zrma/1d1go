package main

import (
	"fmt"

	"1d1go/boj/p2k/p2400"
)

func main() {
	var i, j, k int
	_, _ = fmt.Scan(&i, &j, &k)

	got := p2400.Solve2480([3]int{i, j, k})
	fmt.Println(got)
}
