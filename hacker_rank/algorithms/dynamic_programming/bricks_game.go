package dynamic_programming

import "github.com/zrma/1d1c/hacker_rank/common/utils"

// Use the cache for dynamic programming.
// If not do that, you can not solve it within the time limit.
type cache map[int]map[int]int64

func playSubGame(arr []int32, begin, end int, c cache) int64 {
	if endData, ok := c[begin]; ok {
		if result, ok := endData[end]; ok {
			return result
		}
	} else {
		c[begin] = make(map[int]int64)
	}

	if end-begin <= 3 {
		var sum int64
		for i := 0; begin+i <= end; i++ {
			if i == 3 {
				break
			}
			sum += int64(arr[begin+i])
		}

		c[begin][end] = sum
		return sum
	}

	if end-begin == 4 {
		result := utils.MaxInt64([]int64{
			int64(arr[begin] + arr[begin+1] + arr[begin+2]),
			int64(arr[begin] + arr[begin+4]),
		})
		c[begin][end] = result
		return result
	}

	// I will make the most of my score.
	result := utils.MaxInt64([]int64{
		// The opponent will make the least of my score.
		utils.MinInt64([]int64{
			int64(arr[begin]) + playSubGame(arr, begin+2, end, c),
			int64(arr[begin]) + playSubGame(arr, begin+3, end, c),
			int64(arr[begin]) + playSubGame(arr, begin+4, end, c),
		}),
		// The opponent will make the least of my score.
		utils.MinInt64([]int64{
			int64(arr[begin]) + int64(arr[begin+1]) + playSubGame(arr, begin+3, end, c),
			int64(arr[begin]) + int64(arr[begin+1]) + playSubGame(arr, begin+4, end, c),
			int64(arr[begin]) + int64(arr[begin+1]) + playSubGame(arr, begin+5, end, c),
		}),
		// The opponent will make the least of my score.
		utils.MinInt64([]int64{
			int64(arr[begin]) + int64(arr[begin+1]) + int64(arr[begin+2]) + playSubGame(arr, begin+4, end, c),
			int64(arr[begin]) + int64(arr[begin+1]) + int64(arr[begin+2]) + playSubGame(arr, begin+5, end, c),
			int64(arr[begin]) + int64(arr[begin+1]) + int64(arr[begin+2]) + playSubGame(arr, begin+6, end, c),
		}),
	})
	c[begin][end] = result
	return result
}

func bricksGame(arr []int32) int64 {
	c := make(cache)
	return playSubGame(arr, 0, len(arr)-1, c)
}
