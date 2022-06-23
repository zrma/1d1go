package p11000

import (
	"fmt"
	"strconv"
)

func Solve11054(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	dp0 := make([]int, n)
	dp1 := make([]int, n)

	for i := range dp0 {
		dp0[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp0[i] < dp0[j]+1 {
				dp0[i] = dp0[j] + 1
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if arr[i] > arr[j] && dp1[i] < dp1[j]+1 {
				dp1[i] = dp1[j] + 1
			}
		}
	}

	max := dp0[0] + dp1[0]
	for i := 1; i < n; i++ {
		if max < dp0[i]+dp1[i] {
			max = dp0[i] + dp1[i]
		}
	}
	_, _ = fmt.Fprint(writer, max)
}
