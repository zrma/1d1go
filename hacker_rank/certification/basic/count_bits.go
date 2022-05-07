package basic

func countBits(n uint32) int32 {
	var res int32
	for n > 0 {
		res += int32(n % 2)
		n /= 2
	}
	return res
}
