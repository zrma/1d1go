package p2700

func Solve2753(year int) int {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return 1
	} else {
		return 0
	}
}
