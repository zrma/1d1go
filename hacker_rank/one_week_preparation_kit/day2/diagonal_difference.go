package day2

func diagonalDifference(arr [][]int32) int32 {
	if len(arr) == 0 {
		return 0
	}
	var sum0, sum1 int32
	lastIndex := len(arr) - 1

	// NOTE - time complexity is O(n)
	for i, row := range arr {
		sum0 += row[i]
		sum1 += row[lastIndex-i]
	}
	return abs(sum0 - sum1)
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
