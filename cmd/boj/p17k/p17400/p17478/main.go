package main

import (
	"bufio"
	"os"

	"1d1go/boj/p17k/p17400"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p17400.Solve17478(scanner, writer)
}
