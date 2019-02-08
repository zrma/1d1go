package warm_up

func sockMerchant(ar []int32) int32 {
	counter := make(map[int32]int32)

	for _, n := range ar {
		if _, ok := counter[n]; ok {
			counter[n]++
		} else {
			counter[n] = 1
		}
	}

	var result int32
	for _, v := range counter {
		result += v / 2
	}

	return result
}
