package main

import (
	"bufio"
	"os"

	"1d1go/boj/p10k/p10700"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	p10700.Solve10718(writer)
}
