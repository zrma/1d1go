package p14600

func Solve14681(x, y int) int {
	if x == 0 || y == 0 {
		return -1
	}
	if x > 0 {
		if y > 0 {
			return 1
		} else {
			return 4
		}
	} else {
		if y > 0 {
			return 2
		} else {
			return 3
		}
	}
}
