package strings

func max(i, j int32) int32 {
	if i > j {
		return i
	}

	return j
}

// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func commonChild(s1 string, s2 string) int32 {
	l1 := len(s1) + 1
	l2 := len(s2) + 1
	var arr [][]int32
	for i := 0; i < l2; i++ {
		arr = append(arr, make([]int32, l1))
	}

	for i := 0; i < l1; i++ {
		arr[0][i] = 0
		arr[i][0] = 0
	}

	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				arr[i+1][j+1] = arr[i][j] + 1
			} else {
				arr[i+1][j+1] = max(arr[i+1][j], arr[i][j+1])
			}
		}
	}

	return arr[l1-1][l2-1]
}
