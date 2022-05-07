package p1400

import (
	"1d1go/utils/integer"
)

func Solve1463(n int) int {
	arr := make([]int, n+1)
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + 1
		if i%2 == 0 {
			arr[i] = integer.MinInt(arr[i], arr[i/2]+1)
		}
		if i%3 == 0 {
			arr[i] = integer.MinInt(arr[i], arr[i/3]+1)
		}
	}
	return arr[n]
}
