package day1

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	//goland:noinspection SpellCheckingInspection
	meridiem := s[len(s)-2:]
	hour := s[:2]

	if meridiem == "PM" && hour != "12" {
		hour = strconv.Itoa(12 + int(hour[0]-'0')*10 + int(hour[1]-'0'))
	} else if meridiem == "AM" && hour == "12" {
		hour = "00"
	}
	return fmt.Sprintf("%s%s", hour, s[2:8])
}
