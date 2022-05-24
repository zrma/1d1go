package p2800

import (
	"fmt"
	"strconv"
)

func Solve2869(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	v, _ := strconv.Atoi(scanner.Text())

	if v <= a {
		_, _ = fmt.Fprint(writer, 1)
		return
	}

	res := (v-b-1)/(a-b) + 1
	_, _ = fmt.Fprint(writer, res)
}
