package p2500

func Solve2525(hour, minute, duration int) (int, int) {
	minute += duration
	if minute >= 60 {
		hour += minute / 60
	}
	minute %= 60
	hour %= 24
	return hour, minute
}
