package p2900

import (
	"fmt"
	"strconv"
)

func Solve2908(reader Reader, writer Writer) {
	var s0, s1 string
	_, _ = fmt.Fscan(reader, &s0, &s1)
	a, _ := strconv.Atoi(revStr(s0))
	b, _ := strconv.Atoi(revStr(s1))

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
