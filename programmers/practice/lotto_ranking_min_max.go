package practice

func lottoRankingMinMax(lottoNumbers []int, winNumbers []int) []int {
	var unknownSlots int
	var matchedCount int
	for _, lottoNum := range lottoNumbers {
		if lottoNum == 0 {
			unknownSlots++
		} else {
			for _, winNum := range winNumbers {
				if lottoNum == winNum {
					matchedCount++
				}
			}
		}
	}
	result := min(6, 6-(matchedCount-1))
	return []int{
		max(1, result-unknownSlots),
		result,
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
