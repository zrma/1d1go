package p24400

import (
	"fmt"
	"strconv"
)

func Solve24416(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	res1 := countFibRecursive(n)
	res2 := countFibLoop(n)

	_, _ = fmt.Fprintf(writer, "%d %d", res1, res2)
}

func countFibRecursive(n int) int {
	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func countFibLoop(n int) int {
	// 1, 2 제외
	const countToSkip = 2
	return n - countToSkip
}
