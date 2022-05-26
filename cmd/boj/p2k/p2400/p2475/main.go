package main

import (
	"bufio"
	"os"

	"1d1go/boj/p2k/p2400"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p2400.Solve2475(scanner, writer)
}
