package p1600

import (
	"fmt"
	"strconv"
)

func Solve1620(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	indexes := make(map[string]int)
	names := make([]string, n)

	for i := range names {
		var name string
		_, _ = fmt.Fscan(reader, &name)
		names[i] = name
		indexes[name] = i + 1
	}

	for i := 0; i < m; i++ {
		var name string
		_, _ = fmt.Fscan(reader, &name)

		if _, ok := indexes[name]; ok {
			_, _ = fmt.Fprintln(writer, indexes[name])
		} else {
			idx, _ := strconv.Atoi(name)
			_, _ = fmt.Fprintln(writer, names[idx-1])
		}
	}
}
