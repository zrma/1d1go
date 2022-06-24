package integer

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func LongestIncSubSeqLenSquare[T constraints.Integer](arr []T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}

	dp := make([]T, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	for j := 1; j < len(arr); j++ {
		for i := 0; i < j; i++ {
			if arr[i] < arr[j] {
				dp[j] = Max(dp[i]+1, dp[j])
			}
		}
	}

	return Max(dp...)
}

func LongestIncSubSeqLenLog[T constraints.Integer](arr []T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}

	dp := make([]T, len(arr))
	dp[0] = arr[0]

	i := 0
	for j := 1; j < len(arr); j++ {
		if dp[i] < arr[j] {
			dp[i+1] = arr[j]
			i++
		} else {
			idx := sort.Search(i, func(k int) bool {
				return arr[j] <= arr[k]
			})
			dp[idx] = arr[j]
		}
	}
	return T(i + 1)
}
