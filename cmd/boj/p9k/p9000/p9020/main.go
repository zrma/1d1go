package main

import (
	"bufio"
	"os"

	"1d1go/boj/p9k/p9000"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p9000.Solve9020(scanner, writer)
}
