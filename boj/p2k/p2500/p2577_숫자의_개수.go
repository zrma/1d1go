package p2500

import (
	"fmt"
	"io"
)

func Solve2577(reader io.Reader, writer io.Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	res := [10]int{}
	mul := a * b * c
	for mul > 0 {
		res[mul%10]++
		mul /= 10
	}

	for _, n := range res {
		_, _ = fmt.Fprintln(writer, n)
	}
}
