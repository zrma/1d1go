package warm_up

func jumpingOnClouds(c []int32) int32 {
	var total int32
	var before int32
	l := len(c)
	for cur, v := range c {
		if v == 1 {
			jump := getJump(before, int32(cur-1))
			if jump < 0 {
				return -1
			}

			total += jump + 1
			before = int32(cur) + 1
		} else if cur == (l - 1) {
			total += (int32(cur) - before + 1) / 2
		}
	}

	return total
}

func getJump(before, after int32) int32 {
	diff := after - before
	if diff < 0 {
		return -1
	}

	return (diff + 1) / 2
}
