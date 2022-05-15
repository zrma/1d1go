package main

import (
	"bufio"
	"os"

	"1d1go/boj/p8k/p8900"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p8900.Solve8958(scanner, writer)
}
