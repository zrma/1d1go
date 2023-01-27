package warmup

import (
	"fmt"
	"math"
)

var fmtPrintf = fmt.Printf

func miniMaxSum(arr []int64) {
	var min, max, sum int64
	min = math.MaxInt64
	max = math.MinInt64

	for _, n := range arr {
		if min > n {
			min = n
		}

		if max < n {
			max = n
		}

		sum += n
	}

	_, _ = fmtPrintf("%d %d", sum-max, sum-min)
}
