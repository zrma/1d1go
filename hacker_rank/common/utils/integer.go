package utils

import "math"

func MinInt32(arr []int32) int32 {
	var min int32 = math.MaxInt32
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

func MinInt64(arr []int64) int64 {
	var min int64 = math.MaxInt64
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

func MaxInt32(arr []int32) int32 {
	var max int32 = math.MinInt32
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

func MaxInt64(arr []int64) int64 {
	var max int64 = math.MinInt64
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

func PowInt64(n, p int64) int64 {
	if n == 0 {
		return 0
	}

	var result int64 = 1
	for ; p > 0; p-- {
		result *= n
	}

	return result
}
