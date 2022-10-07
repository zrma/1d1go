package p14000

import (
	"fmt"
	"sort"
)

func Solve14003(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
		dp[i] = 1
	}

	subSeq := make([]int, 0, n)
	subSeq = append(subSeq, arr[0])
	for i := 1; i < n; i++ {
		v := arr[i]

		idx := sort.SearchInts(subSeq, v)
		if idx == len(subSeq) {
			subSeq = append(subSeq, v)
			dp[i] = len(subSeq)
		} else {
			subSeq[idx] = v
			dp[i] = idx + 1
		}
	}

	maxLen := len(subSeq)
	_, _ = fmt.Fprintln(writer, maxLen)

	res := make([]int, maxLen)
	for i := n - 1; i >= 0; i-- {
		if dp[i] == maxLen {
			res[maxLen-1] = arr[i]
			maxLen--
		}
	}

	for i := 0; i < len(res); i++ {
		_, _ = fmt.Fprintf(writer, "%d ", res[i])
	}
}
