package p2100

import (
	"fmt"
	"io"
	"sort"
)

func Solve2141(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	tot := int64(0)
	type village struct {
		x, p int64
	}
	villages := make([]village, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &villages[i].x, &villages[i].p)
		tot += villages[i].p
	}

	sort.Slice(villages, func(i, j int) bool {
		return villages[i].x < villages[j].x
	})

	target := tot / 2
	if tot%2 == 1 {
		target++
	}
	sum := int64(0)
	pos := int64(0)
	for _, v := range villages {
		sum += v.p
		if sum >= target {
			pos = v.x
			break
		}
	}
	_, _ = fmt.Fprint(writer, pos)
}
