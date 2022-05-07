package warmup

func sockMerchant(arr []int32) int32 {
	counter := make(map[int32]int32)

	for _, n := range arr {
		counter[n]++
	}

	var result int32
	for _, v := range counter {
		result += v / 2
	}

	return result
}
