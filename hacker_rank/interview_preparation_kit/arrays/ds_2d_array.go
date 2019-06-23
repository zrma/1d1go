package arrays

import "math"

// HourGlassSum 함수는 주어진 2차원 int32 슬라이스 내부를 순회하며 모래시계 모양의 좌표 내의 값을 합산하고
// 그 결과 중 가장 큰 값을 반환한다.
// 모래 시계 모양의 좌표값 합산 예시는 아래와 같다.
//
// 1 2 3
// 4 5 6
// 7 8 9
//
// 1 + 2 + 3 + 5 + 7 + 8 + 9 = 35
func HourGlassSum(arr [][]int32) int32 {

	rowSize := len(arr)
	if rowSize < 3 {
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
