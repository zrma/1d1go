package p20000

import (
	"fmt"
	"io"

	"1d1go/utils/ds"
)

func Solve20040(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	uf := ds.NewUnionFind(n)

	var res int
	for i := 1; i <= m; i++ {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)

		if uf.Find(a) == uf.Find(b) && res == 0 {
			res = i
		}
		uf.Union(a, b)
	}
	_, _ = fmt.Fprint(writer, res)
}
