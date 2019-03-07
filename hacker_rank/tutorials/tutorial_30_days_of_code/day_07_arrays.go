package tutorial_30_days_of_code

import "fmt"

func printReverse(arr []int32) {
	for i := len(arr); i > 0; i-- {
		fmt.Printf("%d", arr[i-1])
		if i > 1 {
			fmt.Printf(" ")
		}
	}

	fmt.Printf("\n")
}
