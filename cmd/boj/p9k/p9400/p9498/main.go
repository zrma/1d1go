package main

import (
	"fmt"

	"1d1go/boj/p9k/p9400"
)

func main() {
	var score int
	_, _ = fmt.Scan(&score)

	res := p9400.Solve9498(score)
	fmt.Println(res)
}
