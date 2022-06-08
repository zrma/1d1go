package warmup

import (
	"math"
)

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

	_, _ = funcPrintf("%d %d", sum-max, sum-min)
}
