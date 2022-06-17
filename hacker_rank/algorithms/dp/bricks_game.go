package dp

import (
	"1d1go/utils/integer"
)

func bricksGame(arr []int32) int64 {
	length := len(arr)
	if length <= 3 {
		var sum int64
		for _, n := range arr {
			sum += int64(n)
		}
		return sum
	}

	result := make([]int64, length)
	var total int64

	lastIndex := length - 1

	// case 0-2
	// arr    :  1,   2,   3,   4,   5
	// total  : 15 ← 14 ← 12  ← 9  ← 5
	// result :  6,   9,  12,   9,   5

	// case 1
	// arr    :  999,   1,   1,   1,   0
	// total  : 1002 ←  3 ←  2  ← 1  ← 0
	// result : 1001,   2,   2,   1,   0

	// case 3
	// arr    :    0,     1,     1,     1,   999,  999
	// total  : 2001 ← 2001 ← 2000 ← 1999 ← 1998 ← 999
	// result : 1001,  1000,  1001,  1999,  1998,  999

	for i := 0; i < 3; i++ {
		total += int64(arr[lastIndex-i])
		result[lastIndex-i] = total
	}

	for i := 3; i <= lastIndex; i++ {
		total += int64(arr[lastIndex-i])

		// 상대의 몫이 최소가 되도록
		result[lastIndex-i] = total -
			integer.Min(
				result[lastIndex-i+1],
				result[lastIndex-i+2],
				result[lastIndex-i+3],
			)
	}
	return result[0]
}

func bricksGameV2(arr []int32) int64 {
	length := len(arr)
	if length <= 3 {
		var sum int64
		for _, n := range arr {
			sum += int64(n)
		}
		return sum
	}

	lastIndex := length - 1

	tot := make([]int64, length)
	var prev int64
	for i := 0; i < length; i++ {
		n := int64(arr[lastIndex-i])
		cur := prev + n
		prev = cur
		tot[i] = cur
	}

	pick := make([]int64, length)
	for i := 0; i < 3; i++ {
		pick[i] = tot[i]
	}

	getPickNth := func(i int) int64 {
		return tot[i] - pick[i]
	}

	for i := 3; i <= lastIndex; i++ {
		take1 := int64(arr[lastIndex-i]) + getPickNth(i-1)
		take2 := int64(arr[lastIndex-i]+arr[lastIndex-i+1]) + getPickNth(i-2)
		take3 := int64(arr[lastIndex-i]+arr[lastIndex-i+1]+arr[lastIndex-i+2]) + getPickNth(i-3)

		pick[i] = integer.Max(take1, take2, take3)
	}
	return pick[lastIndex]
}

func bricksGameReversal(arr []int32) int64 {
	var sum int64
	for _, n := range arr {
		sum += int64(n)
	}
	return sum - bricksGame(arr)
}

// 메모이제이션을 위해 캐시를 사용한다.
// 그러지 않으면 성능 테스트 케이스 타임아웃에 걸린다.
type brickCache map[int]innerCache

type innerCache map[int]int64

func (cache brickCache) get(begin, end int) (int64, bool) {
	if inner, ok := cache[begin]; ok {
		result, ok := inner[end]
		return result, ok
	}

	cache[begin] = make(innerCache)
	return 0, false
}

func playGameRecur(arr []int32, begin, end int, cache brickCache) int64 {
	if result, ok := cache.get(begin, end); ok {
		return result
	}

	length := end - begin
	// 재귀 함수 탈출 조건
	if length <= 3 {
		var sum int64
		for i := 0; i <= integer.Min(length, 2); i++ {
			sum += int64(arr[begin+i])
		}

		cache[begin][end] = sum
		return sum
	}

	// 나는 최대한 점수를 많이 내려고 한다.
	result := integer.Max(
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		integer.Min(
			int64(arr[begin])+playGameRecur(arr, begin+2, end, cache),
			int64(arr[begin])+playGameRecur(arr, begin+3, end, cache),
			int64(arr[begin])+playGameRecur(arr, begin+4, end, cache),
		),
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		integer.Min(
			int64(arr[begin])+int64(arr[begin+1])+playGameRecur(arr, begin+3, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+playGameRecur(arr, begin+4, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+playGameRecur(arr, begin+5, end, cache),
		),
		// 상대방은 내 점수를 최대한 적게 내도록 한다.
		integer.Min(
			int64(arr[begin])+int64(arr[begin+1])+int64(arr[begin+2])+playGameRecur(arr, begin+4, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+int64(arr[begin+2])+playGameRecur(arr, begin+5, end, cache),
			int64(arr[begin])+int64(arr[begin+1])+int64(arr[begin+2])+playGameRecur(arr, begin+6, end, cache),
		),
	)
	cache[begin][end] = result
	return result
}

func bricksGameRecur(arr []int32) int64 {
	cache := make(brickCache)
	return playGameRecur(arr, 0, len(arr)-1, cache)
}
