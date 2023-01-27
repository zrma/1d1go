package p2100

import (
	"fmt"
	"io"
	"math"
	"sort"
)

func Solve2108(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Ints(arr)

	const (
		offset = 4000
		maxLen = offset * 2
	)
	counts := [maxLen + 1]int{}

	sum := 0.0
	min := math.MaxInt32
	max := math.MinInt32
	for _, v := range arr {
		sum += float64(v)
		counts[v+offset]++
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	avg := int(math.Round(sum / float64(len(arr))))
	_, _ = fmt.Fprintln(writer, avg)

	median := arr[len(arr)/2]
	_, _ = fmt.Fprintln(writer, median)

	haveToRecord2nd := false
	maxCountIdx := 0
	maxCountVal := 0
	for i, cnt := range counts {
		if cnt == maxCountVal && haveToRecord2nd {
			haveToRecord2nd = false
			maxCountIdx = i - offset
		}
		if cnt > maxCountVal {
			haveToRecord2nd = true
			maxCountVal = cnt
			maxCountIdx = i - offset
		}
	}
	_, _ = fmt.Fprintln(writer, maxCountIdx)

	minMaxDiff := max - min
	_, _ = fmt.Fprintln(writer, minMaxDiff)
}
