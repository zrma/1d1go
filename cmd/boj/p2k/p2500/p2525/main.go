package main

import (
	"fmt"

	"1d1go/boj/p2k/p2500"
)

func main() {
	var hour, minute int
	var duration int
	_, _ = fmt.Scan(&hour, &minute)
	_, _ = fmt.Scan(&duration)

	hour, minute = p2500.Solve2525(hour, minute, duration)
	fmt.Println(hour, minute)
}
