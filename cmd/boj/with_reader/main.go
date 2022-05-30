package main

import (
	"bufio"
	"os"

	"1d1go/boj/p1k/p1100"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	// NOTE - run solution to test in cli or IDE
	//        pO000.SolveOOOO(reader, writer)
	p1100.Solve1152(reader, writer)
}
