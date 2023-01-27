package p4300

import (
	"fmt"
	"io"
)

func Solve4344(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		countOverAvg(reader, writer)
	}
}

func countOverAvg(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	sum := 0
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
		sum += arr[i]
	}

	avg := float64(sum) / float64(n)
	count := 0
	for _, v := range arr {
		if float64(v) > avg {
			count++
		}
	}

	_, _ = fmt.Fprintf(writer, "%.3f%%\n", float64(count*100)/float64(n))
}
