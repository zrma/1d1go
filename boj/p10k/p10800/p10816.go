package p10800

import (
	"fmt"
	"strconv"
)

func Solve10816(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		counts[v]++
	}

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < m; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		_, _ = fmt.Fprint(writer, counts[v], " ")
	}
}
