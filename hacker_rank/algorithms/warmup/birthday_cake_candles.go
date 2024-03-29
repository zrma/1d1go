package warmup

import "math"

func birthdayCakeCandles(arr []int32) int32 {
	var cnt, max int32
	max = math.MinInt32

	for _, n := range arr {
		if n > max {
			cnt = 1
			max = n
			continue
		}

		if n == max {
			cnt++
		}
	}

	return cnt
}
