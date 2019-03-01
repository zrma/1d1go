package tutorial_30_days_of_code

import (
	"fmt"
	"strconv"
)

func stringToInteger(s string) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Bad String")
		return
	}

	fmt.Println(num)
}

func StringToInteger(s string) {
	stringToInteger(s)
}
