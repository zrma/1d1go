package tutorial30daysofcode

func bitwiseAND(n, k int32) int32 {
	var max int32
	var i, j int32
	for i = 1; i < n; i++ {
		for j = i + 1; j <= n; j++ {
			and := i & j
			if and < k && max < and {
				max = and
			}
		}
	}
	return max
}
