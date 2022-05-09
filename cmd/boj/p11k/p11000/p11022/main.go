package main

import (
	"bufio"
	"os"
	"strconv"

	"1d1go/boj/p11k/p11000"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr2D := make([][2]int, n)
	for i := range arr2D {
		_ = scanner.Scan()
		arr2D[i][0], _ = strconv.Atoi(scanner.Text())
		_ = scanner.Scan()
		arr2D[i][1], _ = strconv.Atoi(scanner.Text())
	}

	p11000.Solve11022(arr2D)
}
