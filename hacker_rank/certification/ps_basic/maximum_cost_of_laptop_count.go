package ps_basic

func maxCost(cost []int32, labels []string, dailyCount int32) int32 {
	var max, cur int32
	var today int32

	for i, c := range cost {
		cur += c
		if labels[i] == "legal" {
			today++
			if today == dailyCount {
				if max < cur {
					max = cur
				}
				today = 0
				cur = 0
			}
		}
	}
	return max
}
