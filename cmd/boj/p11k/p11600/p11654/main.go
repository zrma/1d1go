package main

import (
	"bufio"
	"os"

	"1d1go/boj/p11k/p11600"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p11600.Solve11654(scanner, writer)
}
