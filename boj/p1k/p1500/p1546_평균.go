package p1500

import (
	"fmt"
	"io"
	"math"
)

func Solve1546(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	max := math.MinInt
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
		if arr[i] > max {
			max = arr[i]
		}
	}
	sum := 0.0
	for i := range arr {
		sum += (float64(arr[i]) / float64(max)) * 100
	}
	_, _ = fmt.Fprint(writer, sum/float64(n))
}
