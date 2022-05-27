package main

import (
	"bufio"
	"os"

	"1d1go/boj/p10k/p10700"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p10700.Solve10757(scanner, writer)
}
