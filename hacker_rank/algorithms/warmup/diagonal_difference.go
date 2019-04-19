package warmup

func diagonalDifference(arr [][]int32) int32 {
	var val1, val2, pos1, pos2 int32
	pos2 = int32(len(arr[0]) - 1)

	for _, row := range arr {
		val1 += row[pos1]
		pos1++

		val2 += row[pos2]
		pos2--
	}

	abs := func(in int32) int32 {
		if in < 0 {
			return -in
		}

		return in
	}

	return abs(val1 - val2)
}
