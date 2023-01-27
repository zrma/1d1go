package p2800

import (
	"fmt"
	"io"
)

func Solve2884(reader io.Reader, writer io.Writer) {
	var hour, minute int
	_, _ = fmt.Fscan(reader, &hour, &minute)

	hour += 23
	minute += 15

	if minute >= 60 {
		minute -= 60
		hour++
	}
	hour %= 24
	_, _ = fmt.Fprintf(writer, "%d %d", hour, minute)
}
