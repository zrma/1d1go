package p3000

import (
	"fmt"
	"strconv"
)

func Solve3052(scanner Scanner, write Writer) {
	m := make(map[int]bool)
	for i := 0; i < 10; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		m[n%42] = true
	}
	_, _ = fmt.Fprint(write, len(m))
}
