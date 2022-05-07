package main

import (
	"fmt"

	"1d1go/boj/p2k/p2800"
)

func main() {
	var hour, minute int
	_, _ = fmt.Scan(&hour, &minute)

	hour, minute = p2800.Solve2884(hour, minute)
	fmt.Println(hour, minute)
}
