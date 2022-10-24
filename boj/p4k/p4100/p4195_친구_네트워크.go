package p4100

import (
	"fmt"

	"1d1go/utils/ds"
)

func Solve4195(reader Reader, writer Writer) {
	var tc int
	_, _ = fmt.Fscan(reader, &tc)

	for i := 0; i < tc; i++ {
		solve4195(reader, writer)
	}
}

func solve4195(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	uf := ds.NewUnionFind(2 * n)
	friends := make(map[string]int)
	idx := 1

	for i := 0; i < n; i++ {
		var a, b string
		_, _ = fmt.Fscan(reader, &a, &b)

		if _, ok := friends[a]; !ok {
			friends[a] = idx
			idx++
		}
		if _, ok := friends[b]; !ok {
			friends[b] = idx
			idx++
		}

		uf.Union(friends[a], friends[b])

		_, _ = fmt.Fprintln(writer, uf.Size(friends[a]))
	}
}
