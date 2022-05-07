package p2800

func Solve2884(hour, minute int) (int, int) {
	hour += 23
	minute += 15

	if minute >= 60 {
		minute -= 60
		hour++
	}
	hour %= 24
	return hour, minute
}
