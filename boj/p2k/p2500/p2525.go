package p2500

func Solve2525(fmt InOut) {
	var hour, minute int
	var duration int
	_, _ = fmt.Scan(&hour, &minute)
	_, _ = fmt.Scan(&duration)

	minute += duration
	if minute >= 60 {
		hour += minute / 60
	}
	minute %= 60
	hour %= 24
	_, _ = fmt.Println(hour, minute)
}
