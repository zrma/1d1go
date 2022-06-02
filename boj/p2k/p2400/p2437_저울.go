package p2400

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve2437(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
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
