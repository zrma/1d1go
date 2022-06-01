package p2800

import (
	"fmt"
	"strconv"
)

func Solve2884(scanner Scanner, writer Writer) {
	scanner.Scan()
	hour, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	minute, _ := strconv.Atoi(scanner.Text())

	hour += 23
	minute += 15

	if minute >= 60 {
		minute -= 60
		hour++
	}
	hour %= 24
	_, _ = fmt.Fprintf(writer, "%d %d", hour, minute)
}
