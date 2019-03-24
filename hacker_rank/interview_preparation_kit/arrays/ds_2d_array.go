package arrays

import "math"

func HourGlassSum(arr [][]int32) int32 {

	rowSize := len(arr)
	if rowSize == 0 || rowSize < 3 {
		return 0
	}

	colSize := len(arr[0])
	if colSize < 3 {
		return 0
	}

	calc := func(x, y int) int32 {
		sum := arr[y-1][x-1] + arr[y-1][x] + arr[y-1][x+1]
		sum += arr[y][x]
		sum += arr[y+1][x-1] + arr[y+1][x] + arr[y+1][x+1]

		return sum
	}

	var max int32 = math.MinInt32
	for y := 1; y < rowSize-1; y++ {
		for x := 1; x < colSize-1; x++ {
			val := calc(x, y)
			if max < val {
				max = val
			}
		}
	}

	return max
}
