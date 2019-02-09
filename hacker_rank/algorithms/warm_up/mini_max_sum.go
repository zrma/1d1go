package warm_up

import (
	"fmt"
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

	fmt.Printf("%d %d", sum-max, sum-min)
}

func MiniMaxSum(arr []int64) {
	miniMaxSum(arr)
}
