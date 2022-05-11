package main

import (
	"bufio"
	"fmt"
	"os"

	"1d1go/boj/p2k/p2700"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	res := p2700.Solve2742(n)
	for _, v := range res {
		_, _ = fmt.Fprintln(w, v)
	}
}
