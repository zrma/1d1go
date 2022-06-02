package p10900

import (
	"fmt"
	"strconv"
)

func Solve10998(scanner Scanner, writer Writer) {
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	_, _ = fmt.Fprint(writer, a*b)
}
