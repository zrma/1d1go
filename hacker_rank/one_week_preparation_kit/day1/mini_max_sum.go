package day1

import (
	"math"
)

func miniMaxSum(arr []int32) {
	var min int64 = math.MaxInt64
	var max int64 = math.MinInt64
	var sum int64 = 0
	for _, n := range arr {
		v := int64(n)
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	_, _ = funcPrintf("%d %d", sum-max, sum-min)
}
