package p14400

import (
	"fmt"
	"strconv"
)

func Solve14425(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	exist := make(map[string]bool)
	for i := 0; i < n; i++ {
		scanner.Scan()
		exist[scanner.Text()] = true
	}

	cnt := 0
	for i := 0; i < m; i++ {
		scanner.Scan()
		if exist[scanner.Text()] {
			cnt++
		}
	}
	_, _ = fmt.Fprint(writer, cnt)
}
