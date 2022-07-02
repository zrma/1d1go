package p2500

import (
	"fmt"
)

func Solve2525(reader Reader, writer Writer) {
	var hour, minute, duration int
	_, _ = fmt.Fscan(reader, &hour, &minute, &duration)

	minute += duration
	if minute >= 60 {
		hour += minute / 60
	}
	minute %= 60
	hour %= 24

	_, _ = fmt.Fprintf(writer, "%d %d", hour, minute)
}
