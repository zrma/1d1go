package p1200

import (
	"fmt"
	"strconv"
)

func Solve1269(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		set[v] = true
	}

	unionCount := 0
	for i := 0; i < m; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		if _, ok := set[v]; ok {
			unionCount++
		}
	}
	_, _ = fmt.Fprint(writer, n+m-unionCount*2)
}
