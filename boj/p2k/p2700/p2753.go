package p2700

import (
	"fmt"
	"strconv"
)

func Solve2753(scanner Scanner, writer Writer) {
	scanner.Scan()
	year, _ := strconv.Atoi(scanner.Text())

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		_, _ = fmt.Fprint(writer, "1")
	} else {
		_, _ = fmt.Fprint(writer, "0")
	}
}
