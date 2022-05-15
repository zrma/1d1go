package p2500

import (
	"fmt"
	"strconv"
)

func Solve2525(scanner Scanner, writer Writer) {
	scanner.Scan()
	hour, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	minute, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	duration, _ := strconv.Atoi(scanner.Text())

	minute += duration
	if minute >= 60 {
		hour += minute / 60
	}
	minute %= 60
	hour %= 24

	_, _ = fmt.Fprintln(writer, hour, minute)
}
