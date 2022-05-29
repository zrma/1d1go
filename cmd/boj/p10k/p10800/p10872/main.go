package main

import (
	"bufio"
	"os"

	"1d1go/boj/p10k/p10800"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p10800.Solve10872(scanner, writer)
}
