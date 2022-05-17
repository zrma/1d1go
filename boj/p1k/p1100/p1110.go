package p1100

import (
	"fmt"
	"strconv"
)

func Solve1110(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	count := 1
	for x := modify(n); x != n; x = modify(x) {
		count++
	}

	_, _ = fmt.Fprint(writer, count)
}

func modify(n int) int {
	remain := n % 10
	sumOfDigit := (n / 10) + remain
	return remain*10 + sumOfDigit%10
}
