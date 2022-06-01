package p2200

import (
	"fmt"
	"strconv"
)

func Solve2231(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	res := findConstructor(n)
	_, _ = fmt.Fprint(writer, res)
}

func findConstructor(n int) int {
	for i := 1; i < n; i++ {
		if decompose(i) == n {
			return i
		}
	}
	return 0
}

func decompose(n int) int {
	res := n
	for ; n > 0; n /= 10 {
		res += n % 10
	}
	return res
}
