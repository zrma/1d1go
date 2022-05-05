package p2400

import (
	"sort"
)

func Solve2437(arr []int) int {
	sort.Ints(arr)

	sum := 0
	for _, n := range arr {
		if sum+1 < n {
			return sum + 1
		}
		sum += n
	}
	return sum + 1
}
