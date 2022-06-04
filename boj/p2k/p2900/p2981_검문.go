package p2900

import (
	"fmt"
	"sort"
	"strconv"
)

func Solve2981(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	for i := range arr {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	sort.Ints(arr)

	gcdOfDiff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		gcdOfDiff = gcd(gcdOfDiff, arr[i+1]-arr[i])
	}

	var res []int
	for i := 2; i*i <= gcdOfDiff; i++ {
		if gcdOfDiff%i == 0 {
			if i*i == gcdOfDiff {
				res = append(res, i)
			} else {
				res = append(res, i, gcdOfDiff/i)
			}
		}
	}
	res = append(res, gcdOfDiff)

	sort.Ints(res)

	for _, v := range res {
		_, _ = fmt.Fprintf(writer, "%d ", v)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
