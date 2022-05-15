package p2500

import (
	"fmt"
	"math"
	"strconv"
)

func Solve2562(scanner Scanner, writer Writer) {
	const count = 9
	var max = math.MinInt
	var pos int
	for i := 1; i <= count; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		if n > max {
			max = n
			pos = i
		}
	}
	_, _ = fmt.Fprintln(writer, max)
	_, _ = fmt.Fprintln(writer, pos)
}
