package p1700

import (
	"fmt"
	"io"
	"sort"
)

func Solve1764(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	names := make(map[string]bool)
	for i := 0; i < n; i++ {
		var name string
		_, _ = fmt.Fscan(reader, &name)
		names[name] = true
	}

	size := n
	if size < m {
		size = m
	}
	res := make([]string, 0, size)
	for i := 0; i < m; i++ {
		var name string
		_, _ = fmt.Fscan(reader, &name)
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
