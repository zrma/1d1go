package lv0easy

var isBadVersion func(int) bool

func solveFirstBadVersion(maxVer, badVer int) (int, int) {
	var callCount int
	isBadVersion = func(curVer int) bool {
		callCount++
		return curVer >= badVer
	}
	return firstBadVersion(maxVer), callCount
}

func firstBadVersion(maxVer int) int {
	l, r := 1, maxVer
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
