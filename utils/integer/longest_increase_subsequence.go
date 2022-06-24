package integer

import (
	"golang.org/x/exp/constraints"
)

func LongestIncSubSeqLen[T constraints.Integer](arr []T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}

	dp := make([]T, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	return Max(dp...)
}
