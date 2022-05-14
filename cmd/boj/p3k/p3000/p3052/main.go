package main

import (
	"bufio"
	"os"

	"1d1go/boj/p3k/p3000"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p3000.Solve3052(scanner, writer)
}
