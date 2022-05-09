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

	res := p2700.Solve2741(n)
	for _, v := range res {
		_, _ = fmt.Fprintln(w, v)
	}
	_ = w.Flush()
}
