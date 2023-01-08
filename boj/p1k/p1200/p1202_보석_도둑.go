package p1200

import (
	"fmt"
	"sort"

	"1d1go/utils/ds"
)

func Solve1202(reader Reader, writer Writer) {
	var n, k int
	_, _ = fmt.Fscan(reader, &n, &k)

	jewels := make([]jewel, n)
	for i := range jewels {
		_, _ = fmt.Fscan(reader, &jewels[i].weight, &jewels[i].value)
	}

	bags := make([]int, k)
	for i := range bags {
		_, _ = fmt.Fscan(reader, &bags[i])
	}

	_, _ = fmt.Fprint(writer, solve1202(jewels, bags))
}

type jewel struct {
	weight int
	value  int
}

func solve1202(jewels []jewel, bags []int) int {
	sort.Slice(jewels, func(i, j int) bool {
		return jewels[i].weight > jewels[j].weight
	})
	sort.Ints(bags)

	res := 0
	pq := ds.NewMaxHeap(len(bags))
	for _, bag := range bags {
		for len(jewels) > 0 && jewels[len(jewels)-1].weight <= bag {
			pq.Push(jewels[len(jewels)-1].value)
			jewels = jewels[:len(jewels)-1]
		}
		if pq.Size() > 0 {
			v, _ := pq.Pop()
			res += v
		}
	}
	return res
}
