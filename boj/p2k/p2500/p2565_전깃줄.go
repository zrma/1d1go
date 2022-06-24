package p2500

import (
	"fmt"
	"sort"
	"strconv"

	"1d1go/utils/integer"
)

func Solve2565(scanner Scanner, writer Writer) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i][0], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		arr[i][1], _ = strconv.Atoi(scanner.Text())
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	dp := make([]int, n)
	for i := range arr {
		dp[i] = arr[i][1]
	}

	longestIncSubSeqLen := integer.LongestIncSubSeqLenSquare(dp)
	res := n - longestIncSubSeqLen
	_, _ = fmt.Fprint(writer, res)
}
