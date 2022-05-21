package p2900

import (
	"fmt"
	"strconv"
)

func Solve2908(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(revStr(scanner.Text()))
	scanner.Scan()
	b, _ := strconv.Atoi(revStr(scanner.Text()))

	max := a
	if b > max {
		max = b
	}
	_, _ = fmt.Fprint(writer, max)
}

func revStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
