package warmup

func sockMerchant(arr []int32) int32 {
	counts := make(map[int32]int32)

	for _, n := range arr {
		counts[n]++
	}

	var result int32
	for _, v := range counts {
		result += v / 2
	}

	return result
}
