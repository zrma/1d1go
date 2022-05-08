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

	arr := make([]int, n)
	for i := range arr {
		_ = scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	res := p10800.Solve10818(arr)
	fmt.Println(res[0], res[1])
}
