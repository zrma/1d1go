package main

import (
	"bufio"
	"os"
	"strconv"

	"1d1go/boj/p10k/p10900"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	for scanner.Err() == nil {
		if ok := scanner.Scan(); !ok {
			break
		}
		a, _ := strconv.Atoi(scanner.Text())
		if ok := scanner.Scan(); !ok {
			break
		}
		b, _ := strconv.Atoi(scanner.Text())
		if a == 0 && b == 0 {
			break
		}
		p10900.Solve10951(w, a, b)
		_ = w.Flush()
	}
}
