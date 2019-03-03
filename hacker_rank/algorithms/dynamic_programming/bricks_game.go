package dynamic_programming

import "github.com/zrma/1d1c/hacker_rank/common/utils"

// Use the cache for memoize.
// 메모이제이션을 위해 캐시를 사용한다.
// If not do that, you can not solve it within the time limit.
// 그러지 않을 경우, 제한 시간 안에 문제를 풀 수 없다.
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
	// 나는 최대한 점수를 많이 내려고 한다.
	result := utils.MaxInt64([]int64{
		// The opponent will make the least of my score.
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		utils.MinInt64([]int64{
			int64(arr[begin]) + playSubGame(arr, begin+2, end, c),
			int64(arr[begin]) + playSubGame(arr, begin+3, end, c),
			int64(arr[begin]) + playSubGame(arr, begin+4, end, c),
		}),
		// The opponent will make the least of my score.
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		utils.MinInt64([]int64{
			int64(arr[begin]) + int64(arr[begin+1]) + playSubGame(arr, begin+3, end, c),
			int64(arr[begin]) + int64(arr[begin+1]) + playSubGame(arr, begin+4, end, c),
			int64(arr[begin]) + int64(arr[begin+1]) + playSubGame(arr, begin+5, end, c),
		}),
		// The opponent will make the least of my score.
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
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
