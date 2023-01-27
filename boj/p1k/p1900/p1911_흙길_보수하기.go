package p1900

import (
	"fmt"
	"io"
	"sort"
)

func Solve1911(reader io.Reader, writer io.Writer) {
	var n, l int
	_, _ = fmt.Fscan(reader, &n, &l)

	type hole struct {
		start, end int
	}
	holes := make([]hole, n)
	for i := range holes {
		_, _ = fmt.Fscan(reader, &holes[i].start, &holes[i].end)
	}

	sort.Slice(holes, func(i, j int) bool {
		if holes[i].start == holes[j].start {
			return holes[i].end < holes[j].end
		}
		return holes[i].start < holes[j].start
	})

	count := 0
	covered := 0
	for _, h := range holes {
		if covered < h.start {
			covered = h.start
		}
		for covered < h.end {
			covered += l
			count++
		}
	}
	_, _ = fmt.Fprint(writer, count)
}
