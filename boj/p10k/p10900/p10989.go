package p10900

import (
	"fmt"
	"strconv"
)

func Solve10989(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	res := countingSort(arr)

	for _, v := range res {
		_, _ = fmt.Fprintln(writer, v)
	}
}

func countingSort(arr []int) []int {
	const max = 10000
	count := [max + 1]int{}
	for _, n := range arr {
		count[n]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	res := make([]int, len(arr))
	for _, n := range arr {
		idx := count[n] - 1
		res[idx] = n
		count[n]--
	}
	return res
}
