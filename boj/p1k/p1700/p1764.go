package p1700

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve1764(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	names := make(map[string]bool)
	for i := 0; i < n; i++ {
		scanner.Scan()
		names[scanner.Text()] = true
	}

	size := n
	if size < m {
		size = m
	}
	res := make([]string, 0, size)
	for i := 0; i < m; i++ {
		scanner.Scan()
		name := scanner.Text()
		if _, ok := names[name]; ok {
			res = append(res, name)
		}
	}
	sort.Strings(res)

	_, _ = fmt.Fprintln(writer, len(res))
	for _, name := range res {
		_, _ = fmt.Fprintln(writer, name)
	}
}
