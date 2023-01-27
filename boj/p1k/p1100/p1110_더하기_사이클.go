package p1100

import (
	"fmt"
	"io"
)

func Solve1110(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

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
