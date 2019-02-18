package dictionaries_and_hashmaps

func countTriplets(arr []int64, r int64) int64 {
	if r == 1 {
		map1 := make(map[int64]int64)
		for _, num := range arr {
			map1[num]++
		}

		var total int64
		for _, num := range map1 {
			// length Combinations of 3
			// lC3
			// lC2 = l * (l - 1) * (l - 2) / (3 * 2 * 1)
			total += num * (num - 1) * (num - 2) / 6
		}
		return total
	}

	map2 := make(map[int64]int64)
	map3 := make(map[int64]int64)

	var total int64
	for _, num := range arr {
		map2[num*r]++
		map3[num*r] += map2[num]
		total += map3[num]
	}

	return total
}
