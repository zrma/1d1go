package tutorial30daysofcode

func printReverse(arr []int32) {
	for i := len(arr); i > 0; i-- {
		_, _ = funcPrintf("%d", arr[i-1])
		if i > 1 {
			_, _ = funcPrintf(" ")
		}
	}
}
