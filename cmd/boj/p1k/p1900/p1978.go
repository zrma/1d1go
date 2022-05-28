package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1900"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p1900.Solve1978(scanner, writer)
}
