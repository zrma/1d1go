package day1

import (
	"fmt"
	"math"
)

var fmtPrintf = fmt.Printf

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
	_, _ = fmtPrintf("%d %d", sum-max, sum-min)
}
