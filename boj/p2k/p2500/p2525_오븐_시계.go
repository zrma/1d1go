package p2500

import (
	"fmt"
	"io"
)

func Solve2525(reader io.Reader, writer io.Writer) {
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
