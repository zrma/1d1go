package warm_up

import "math"

func birthdayCakeCandles(ar []int32) int32 {
	var cnt int32
	var max int32
	max = math.MinInt32

	for _, n := range ar {
		if n > max {
			cnt = 1
			max = n
			continue
		}

		if n == max {
			cnt++
			continue
		}
	}

	return cnt
}
