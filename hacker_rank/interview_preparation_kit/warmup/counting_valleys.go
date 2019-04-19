package warmup

func countingValleys(n int32, s string) int32 {
	if n != int32(len(s)) {
		return 0
	}

	var height int
	var cnt int32
	var i int32
	for i = 0; i < n; i++ {
		switch s[i] {
		case 'U':
			if height == -1 {
				cnt++
			}
			height++
		case 'D':
			height--
		default:
			return 0
		}
	}

	return cnt
}
