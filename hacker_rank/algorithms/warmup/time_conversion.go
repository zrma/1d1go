package warmup

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	offset := 0
	if strings.Contains(s, "PM") {
		offset = 12
	}

	token := strings.Split(s[:len(s)-2], ":")
	hour, minute, second := token[0], token[1], token[2]

	hourNum, err := strconv.ParseInt(hour, 10, 64)
	if err != nil {
		hourNum = 0
	}

	return fmt.Sprintf("%02d:%s:%s", int(hourNum)%12+offset, minute, second)
}
