package p2700

import (
	"fmt"
	"io"
)

func Solve2798(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}
	res := findMaxSumOf3Nums(m, arr)
	_, _ = fmt.Fprint(writer, res)
}

func findMaxSumOf3Nums(limit int, arr []int) int {
	max := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				sum := arr[i] + arr[j] + arr[k]
				if sum > max && sum <= limit {
					max = sum
				}
			}
		}
	}
	return max
}
