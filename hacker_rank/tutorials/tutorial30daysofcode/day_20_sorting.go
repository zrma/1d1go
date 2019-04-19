package tutorial30daysofcode

import "fmt"

func sorting(arr []int32) {

	n := len(arr)
	var total int
	for i := 0; i < n; i++ {
		numberOfSwaps := 0

		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numberOfSwaps++
			}
		}

		if numberOfSwaps == 0 {
			break
		}
		total += numberOfSwaps
	}

	fmt.Printf("Array is sorted in %d swaps.\n", total)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[n-1])
}
