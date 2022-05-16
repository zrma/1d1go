package p1000

import (
	"fmt"
	"strconv"
)

func Solve1065(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var count int
	for i := 1; i <= n; i++ {
		if arithmeticProgression(i) {
			count++
		}
	}
	_, _ = fmt.Fprintln(writer, count)
}

// arithmeticProgression args: n (1~1000)
func arithmeticProgression(n int) bool {
	if n < 100 {
		return true
	}
	if n > 999 {
		return false
	}
	i := n / 100
	j := n % 100 / 10
	k := n % 10
	return i-j == j-k
}
