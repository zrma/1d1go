package warmup

func diagonalDifference(arr [][]int32) int32 {
	if len(arr) == 0 {
		return 0
	}

	var val1, val2 int32
	lastIndex := len(arr[0]) - 1

	for pos, row := range arr {
		val1 += row[pos]
		val2 += row[lastIndex-pos]
	}

	abs := func(n int32) int32 {
		if n < 0 {
			return -n
		}

		return n
	}
	return abs(val1 - val2)
}
