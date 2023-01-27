package arrays

import (
	"fmt"
)

type data struct {
	val   int32
	moved int
}

func bubbleSort(items []data, limit int) (int, bool) {
	var (
		n      = len(items)
		cnt    = 0
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i].val > items[i+1].val {
				items[i+1], items[i] = items[i], items[i+1]
				items[i+1].moved++
				cnt++
				if items[i+1].moved > limit {
					return -1, false
				}
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}

	return cnt, true
}

var fmtPrint = fmt.Print

func minimumBribes(arr []int32) {
	adapter := make([]data, len(arr))
	for i, n := range arr {
		adapter[i] = data{n, 0}
	}

	cnt, ok := bubbleSort(adapter, 2)
	if !ok {
		_, _ = fmtPrint("Too chaotic")
		return
	}
	_, _ = fmtPrint(cnt)
}
