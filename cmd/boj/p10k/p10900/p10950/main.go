package main

import (
	"bufio"
	"os"

	"1d1go/boj/p10k/p10900"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p10900.Solve10950(scanner, writer)
}
