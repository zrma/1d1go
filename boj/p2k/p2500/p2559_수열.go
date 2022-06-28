package p2500

import (
	"fmt"
	"strconv"
)

func Solve2559(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	max := 0
	sum := 0
	for i := 0; i < k; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		sum += arr[i]
	}
	max = sum
	for i := k; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())

		sum -= arr[i-k]
		sum += arr[i]
		if sum > max {
			max = sum
		}
	}
	_, _ = fmt.Fprint(writer, max)
}
