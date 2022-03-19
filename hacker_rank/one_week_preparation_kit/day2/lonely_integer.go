package day2

func LonelyInteger(a []int32) int32 {
	var res int32
	// NOTE - time complexity O(n)
	//        if I use map, it will follow golang map implementation's own
	//        if I use sort, it will follow comparison sort O(n log n)
	for i := 0; i < len(a); i++ {
		res ^= a[i]
	}
	return res
}
