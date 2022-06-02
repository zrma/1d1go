package p2700

import (
	"fmt"
	"strconv"
)

func Solve2798(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
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
