package p13300

import (
	"fmt"
	"io"
)

func Solve13305(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	distances := make([]int, n-1)
	for i := range distances {
		_, _ = fmt.Fscan(reader, &distances[i])
	}
	prices := make([]int, n)
	for i := range prices {
		_, _ = fmt.Fscan(reader, &prices[i])
	}

	minPrice := prices[0]
	moved := distances[0]
	res := 0
	for i := 1; i < len(distances); i++ {
		distance := distances[i]
		if curPrice := prices[i]; minPrice > curPrice {
			res += minPrice * moved
			minPrice = curPrice
			moved = distance
		} else {
			moved += distance
		}
	}
	res += minPrice * moved
	_, _ = fmt.Fprint(writer, res)
}
