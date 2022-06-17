package integer

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Integer](arr ...T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	min := arr[0]
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

func Max[T constraints.Integer](arr ...T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	max := arr[0]
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

func Pow[T constraints.Integer](n, p T) T {
	if n == 0 {
		return 0
	}

	var result T = 1
	for ; p > 0; p-- {
		result *= n
	}

	return result
}
