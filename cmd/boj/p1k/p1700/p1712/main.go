package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1700"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p1700.Solve1712(scanner, writer)
}
