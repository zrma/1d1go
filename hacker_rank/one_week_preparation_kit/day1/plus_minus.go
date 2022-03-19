package day1

import "fmt"

func plusMinus(arr []int32) {
	var pos, neg, zero float64
	// NOTE - time complexity O(n)
	for _, n := range arr {
		if n > 0 {
			pos++
		} else if n < 0 {
			neg++
		} else {
			zero++
		}
	}

	const fmtStr = "%.6f\n"
	fmt.Printf(fmtStr, pos/float64(len(arr)))
	fmt.Printf(fmtStr, neg/float64(len(arr)))
	fmt.Printf(fmtStr, zero/float64(len(arr)))
}
