package main

import (
	"bufio"
	"os"

	"1d1go/boj/p15k/p15500"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p15500.Solve15552(scanner, writer)
}
