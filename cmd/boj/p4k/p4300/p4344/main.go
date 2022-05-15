package main

import (
	"bufio"
	"os"

	"1d1go/boj/p4k/p4300"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p4300.Solve4344(scanner, writer)
}
