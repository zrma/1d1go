package practice

func Disguise(clothes [][]string) int {
	counts := make(map[string]int)
	for _, cloth := range clothes {
		part := cloth[1]
		counts[part]++
	}
	result := 1
	for _, count := range counts {
		result *= count + 1
	}
	return result - 1
}
