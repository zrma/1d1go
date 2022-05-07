package sorting

import (
	"fmt"
)

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

	fmt.Printf("Array is sorted in %d swaps.\n", swapCount)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[n-1])
}
