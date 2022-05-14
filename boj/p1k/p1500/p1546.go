package p1500

import (
	"fmt"
	"math"
	"strconv"
)

func Solve1546(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	max := math.MinInt
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		if arr[i] > max {
			max = arr[i]
		}
	}
	sum := 0.0
	for i := range arr {
		sum += (float64(arr[i]) / float64(max)) * 100
	}
	_, _ = fmt.Fprintln(writer, sum/float64(n))
}
