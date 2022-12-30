package p2000

import (
	"fmt"
	"sort"
)

func Solve2012(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	ranks := make([]int, n)
	for i := range ranks {
		_, _ = fmt.Fscan(reader, &ranks[i])
	}

	sort.Ints(ranks)

	res := 0
	for i, r := range ranks {
		diff := i + 1 - r
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}

	_, _ = fmt.Fprint(writer, res)
}
