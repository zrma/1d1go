package p2500

import (
	"math"

	"1d1go/utils"
)

func Solve2562(fmt utils.InOut) {
	const count = 9
	var max = math.MinInt
	var pos int
	for i := 1; i <= count; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		if n > max {
			max = n
			pos = i
		}
	}
	_, _ = fmt.Println(max)
	_, _ = fmt.Println(pos)
}
