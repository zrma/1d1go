package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1500"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	// NOTE - run solution to test in cli or IDE
	//        pO000.SolveOOOO(scanner, writer)
	p1500.Solve1543(scanner, writer)
}
