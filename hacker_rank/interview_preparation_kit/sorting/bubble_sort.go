package sorting

import (
	"fmt"
)

var fmtPrintf = fmt.Printf

func countSwaps(arr []int32) {
	swapCount := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapCount++
			}
		}
	}

	_, _ = fmtPrintf("Array is sorted in %d swaps.\n", swapCount)
	_, _ = fmtPrintf("First Element: %d\n", arr[0])
	_, _ = fmtPrintf("Last Element: %d\n", arr[n-1])
}
