package p2400

import (
	"fmt"
	"io"
	"sort"
)

func Solve2437(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := range arr {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	res := getUnavailableSum(arr)
	_, _ = fmt.Fprint(writer, res)
}

func getUnavailableSum(arr []int) int {
	sort.Ints(arr)

	sum := 0
	for _, n := range arr {
		if sum+1 < n {
			return sum + 1
		}
		sum += n
	}
	return sum + 1
}
