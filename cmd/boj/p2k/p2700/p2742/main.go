package main

import (
	"bufio"
	"os"

	"1d1go/boj/p2k/p2700"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p2700.Solve2742(scanner, writer)
}
