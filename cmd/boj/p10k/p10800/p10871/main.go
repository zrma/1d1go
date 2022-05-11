package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"1d1go/boj/p10k/p10800"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	_ = scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	_ = scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		_ = scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	res := p10800.Solve10871(x, arr)
	for _, v := range res {
		_, _ = fmt.Fprintf(w, "%d ", v)
	}
}
