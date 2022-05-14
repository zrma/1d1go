package p2500

import (
	"fmt"
	"strconv"
)

func Solve2577(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

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
