package day1

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
	_, _ = funcPrintf(fmtStr, pos/float64(len(arr)))
	_, _ = funcPrintf(fmtStr, neg/float64(len(arr)))
	_, _ = funcPrintf(fmtStr, zero/float64(len(arr)))
}
