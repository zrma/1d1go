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
		// 등비수열인 세 숫자 중 첫 번째 숫자가 등장할 경우

		// 계속 순회하다 보면 다음에 언젠가 두 번째 수가 등장할 수도 있다.
		// 그러므로 그 쪽(두 번째)에 미리 비축. increase 1
		map2[num*r]++

		// 이전 이터레이션에서 언젠가 등장한 두 번째 수를 세 번째로 합산 이관(비축).
		map3[num*r] += map2[num]

		// 세 번째 이터레이션이 마무리 되었다면 이전의 누적이 모두 이관 완료 된 것. 정산.
		total += map3[num]
	}

	return total
}
