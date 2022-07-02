package arrays

import "math"

func arrayManipulation(n int, queries [][]int) int64 {
	arr := make([]int64, n+1)
	for _, row := range queries {
		begin, end, value := row[0], row[1], row[2]
		arr[begin-1] += int64(value)
		arr[end] -= int64(value)
	}

	var max int64 = math.MinInt64
	var cur int64
	for _, num := range arr {
		cur += num
		if max < cur {
			max = cur
		}
	}

	return max
}
