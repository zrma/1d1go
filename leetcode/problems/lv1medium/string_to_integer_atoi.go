package lv1medium

import (
	"math"
	"strconv"
	"strings"
)

func myAtoI(str string) int {
	str = strings.TrimLeft(str, " ")
	if str == "" {
		return 0
	}

	var negative bool
	negative, str = parseSign(str)

	str = strings.TrimLeft(str, "0")
	if str == "" {
		return 0
	}

	num := parseDigit(str)
	if negative {
		num = -num
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	if num < math.MinInt32 {
		num = math.MinInt32
	}
	return int(num)
}

func parseSign(str string) (bool, string) {
	var negative bool
	switch str[0] {
	case '-':
		negative = true
		str = str[1:]
	case '+':
		str = str[1:]
	}
	return negative, str
}

func parseDigit(str string) (num int64) {
	for i, c := range str {
		if i > 10 {
			break
		}
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			cur, _ := strconv.Atoi(string(c))
			num = num*10 + int64(cur)
		default:
			return
		}
	}
	return
}
