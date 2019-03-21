package tutorial_30_days_of_code

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func roundTime(input float64) int {
	result := math.Floor(input + 0.5)

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

const (
	day   = 24 * 60 * 60
	month = day * 30
)

func convertCalendar(s string) string {
	ss := strings.Split(s, " ")
	day, month, year := ss[0], ss[1], ss[2]
	return fmt.Sprintf("%s %s %04s", day, month, year)
}

func nestedLogic(from, to time.Time) int {
	sub := from.Sub(to)
	if sub < 0 {
		return 0
	}

	if from.Year() == to.Year() {
		if from.Month() == to.Month() {
			return 15 * roundTime(sub.Seconds()/day)
		} else {
			return 500 * roundTime(sub.Seconds()/month)
		}
	}

	return 10000
}
