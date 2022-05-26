package main

import (
	"bufio"
	"os"

	"1d1go/boj/p10k/p10100"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p10100.Solve10157(scanner, writer)
}
