package main

import (
	"bufio"
	"os"

	"1d1go/boj/p2k/p2500"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p2500.Solve2588(scanner, writer)
}
