package p18108

import (
	"fmt"
	"strconv"
)

func Solve18108(scanner Scanner, writer Writer) {
	scanner.Scan()
	year, _ := strconv.Atoi(scanner.Text())

	res := year - 543
	_, _ = fmt.Fprint(writer, res)
}
