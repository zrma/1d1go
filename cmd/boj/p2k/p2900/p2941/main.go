package main

import (
	"bufio"
	"os"

	"1d1go/boj/p2k/p2900"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p2900.Solve2941(reader, writer)
}
