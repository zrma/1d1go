package main

import (
	"bufio"
	"os"

	"1d1go/boj/p5k/p5600"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p5600.Solve5622(reader, writer)
}
