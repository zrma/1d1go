package tutorial30daysofcode

import (
	"strconv"
)

func stringToInteger(s string) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		_, _ = funcPrintln("Bad String")
		return
	}

	_, _ = funcPrintln(num)
}
