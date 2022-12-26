package p1000

import (
	"fmt"

	"1d1go/utils/ds"
)

func Solve1043(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	uf := ds.NewUnionFind(n)

	var knownCount int
	_, _ = fmt.Fscan(reader, &knownCount)
	for i := 0; i < knownCount; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)

		uf.Union(0, v)
	}

	party := make([][]int, m)
	for i := range party {
		var partyCount int
		_, _ = fmt.Fscan(reader, &partyCount)
		party[i] = make([]int, partyCount)
		for j := range party[i] {
			_, _ = fmt.Fscan(reader, &party[i][j])
			uf.Union(party[i][0], party[i][j])
		}
	}

	if knownCount == 0 {
		_, _ = fmt.Fprint(writer, m)
		return
	}

	res := 0
	for _, p := range party {
		if uf.Find(0) != uf.Find(p[0]) {
			res++
		}
	}
	_, _ = fmt.Fprint(writer, res)
}
