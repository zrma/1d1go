package p2600

import (
	"fmt"
	"sort"

	"1d1go/boj/p2k/p2500"
)

func Solve2667(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	table := make([][]bool, n)
	for y := 0; y < n; y++ {
		table[y] = make([]bool, n)
		var s string
		_, _ = fmt.Fscan(reader, &s)
		for x, c := range s {
			if c == '0' {
				table[y][x] = true
			}
		}
	}

	res := p2500.CountGroupsInTable(table, n, n)
	sort.Ints(res)
	_, _ = fmt.Fprintln(writer, len(res))
	for _, v := range res {
		_, _ = fmt.Fprintln(writer, v)
	}
}
