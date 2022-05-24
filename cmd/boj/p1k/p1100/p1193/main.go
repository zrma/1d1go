package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1100"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p1100.Solve1193(scanner, writer)
}
