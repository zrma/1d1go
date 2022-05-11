package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"1d1go/boj/p15k/p15500"
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

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	res := p15500.Solve15552(arr2D)
	for _, v := range res {
		_, _ = fmt.Fprintln(w, v)
	}
}
