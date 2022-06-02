package p1600

import (
	"fmt"
	"strconv"
)

func Solve1620(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	indexes := make(map[string]int)
	names := make([]string, n)

	for i := range names {
		scanner.Scan()
		name := scanner.Text()
		names[i] = name
		indexes[name] = i + 1
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		name := scanner.Text()
		if _, ok := indexes[name]; ok {
			_, _ = fmt.Fprintln(writer, indexes[name])
		} else {
			idx, _ := strconv.Atoi(name)
			_, _ = fmt.Fprintln(writer, names[idx-1])
		}
	}
}
