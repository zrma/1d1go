package main

import (
	"bufio"
	"os"

	"1d1go/boj/p10k/p10250"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p10250.Solve10250(scanner, writer)
}
