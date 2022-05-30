package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1000"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	// NOTE - run solution to test in cli or IDE
	//        pO000.SolveOOOO(scanner, writer)
	p1000.Solve1000(scanner, writer)
}
