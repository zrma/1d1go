package p10800

import (
	"fmt"
	"math"
	"strconv"
)

func Solve10818(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	_, _ = fmt.Fprintf(writer, "%d %d", min, max)
}
