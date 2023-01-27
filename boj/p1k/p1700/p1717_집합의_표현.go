package p1700

import (
	"fmt"
	"io"

	"1d1go/utils/ds"
)

func Solve1717(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	uf := ds.NewUnionFind(n)

	for i := 1; i <= m; i++ {
		var op, a, b int
		_, _ = fmt.Fscan(reader, &op, &a, &b)

		if op == 0 {
			uf.Union(a, b)
		} else {
			if uf.Find(a) == uf.Find(b) {
				_, _ = fmt.Fprintln(writer, "YES")
			} else {
				_, _ = fmt.Fprintln(writer, "NO")
			}
		}
	}
}
