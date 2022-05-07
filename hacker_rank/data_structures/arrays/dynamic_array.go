package arrays

func dynamicArray(n int32, queries [][]int32) []int32 {
	var lastAnswer int32
	var result []int32
	seqList := make([][]int32, n)

	for _, query := range queries {
		cmd, x, y := query[0], query[1], query[2]

		if cmd == 1 {
			seqIndex := (x ^ lastAnswer) % n
			seq := seqList[seqIndex]
			seq = append(seq, y)
			seqList[seqIndex] = seq
		} else {
			seqIndex := (x ^ lastAnswer) % n
			seq := seqList[seqIndex]
			elemIndex := y % int32(len(seq))
			lastAnswer = seq[elemIndex]
			result = append(result, lastAnswer)
		}
	}
	return result
}
