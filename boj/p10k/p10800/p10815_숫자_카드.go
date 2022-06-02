package p10800

import (
	"fmt"
	"strconv"
)

func Solve10815(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	exist := make(map[int]bool)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		exist[v] = true
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < m; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		if exist[v] {
			_, _ = fmt.Fprint(writer, "1 ")
		} else {
			_, _ = fmt.Fprint(writer, "0 ")
		}
	}
}
