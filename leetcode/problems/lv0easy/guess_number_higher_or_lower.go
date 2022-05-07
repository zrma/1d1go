package lv0easy

var guess func(int) int

func solveGuessNumber(maxNum, target int) (int, int) {
	var callCount int
	guess = func(i int) int {
		callCount++
		switch {
		case target < i:
			return -1
		case target > i:
			return 1
		default:
			return 0
		}
	}
	return guessNumber(maxNum), callCount
}

func guessNumber(maxNum int) int {
	l, r := 1, maxNum
	for l <= r {
		mid1 := l + (r-l)/3
		mid2 := r - (r-l)/3
		v1 := guess(mid1)
		v2 := guess(mid2)
		switch {
		case v1 == 0:
			return mid1
		case v2 == 0:
			return mid2
		case v1 < 0:
			r = mid1 - 1
		case v2 > 0:
			l = mid2 + 1
		default:
			l = mid1 + 1
			r = mid2 - 1
		}
	}
	return -1
}
