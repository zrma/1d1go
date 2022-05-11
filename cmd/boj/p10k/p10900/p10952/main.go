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

	var arr2D [][2]int
	for {
		_ = scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		_ = scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		if a == 0 && b == 0 {
			break
		}
		arr2D = append(arr2D, [2]int{a, b})
	}
	p10900.Solve10952(arr2D)
}
