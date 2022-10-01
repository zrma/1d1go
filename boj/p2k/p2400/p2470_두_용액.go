package p2400

import (
	"fmt"
	"sort"
)

func Solve2470(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	sort.Ints(arr)

	lo, hi := 0, n-1
	loVal, hiVal := arr[lo], arr[hi]
	min := abs(loVal + hiVal)
	for lo < hi {
		sum := arr[lo] + arr[hi]
		if abs(sum) < min {
			min = abs(sum)
			loVal, hiVal = arr[lo], arr[hi]
		}
		if sum < 0 {
			lo++
			continue
		}
		hi--
	}

	_, _ = fmt.Fprint(writer, loVal, hiVal)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
