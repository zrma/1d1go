package tutorial30daysofcode

import (
	"strconv"
)

func binaryNumbers(n int32) {
	s := strconv.FormatInt(int64(n), 2)
	var cnt, max int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '1' {
			cnt++
			if cnt > max {
				max = cnt
			}
		} else {
			cnt = 0
		}
	}

	_, _ = funcPrint(max)
}
