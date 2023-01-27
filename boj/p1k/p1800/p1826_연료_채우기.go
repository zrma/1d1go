package p1800

import (
	"fmt"
	"io"
	"sort"

	"1d1go/utils/ds"
)

func Solve1826(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	type station struct {
		distance, fuel int
	}

	stations := make([]station, n+1)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &stations[i].distance)
		_, _ = fmt.Fscan(reader, &stations[i].fuel)
	}

	var l, p int
	_, _ = fmt.Fscan(reader, &l, &p)

	stations[n] = station{l, 0}

	sort.Slice(stations, func(i, j int) bool {
		return stations[i].distance < stations[j].distance
	})

	h := ds.NewMaxHeap(n)

	res := 0
	for _, curr := range stations {
		if p >= l {
			break
		}

		if p >= curr.distance {
			h.Push(curr.fuel)
			continue
		}

		for p < curr.distance {
			if h.Size() == 0 {
				_, _ = fmt.Fprint(writer, -1)
				return
			}
			v, _ := h.Pop()
			p += v
			res++
		}

		h.Push(curr.fuel)
	}
	_, _ = fmt.Fprint(writer, res)
}
