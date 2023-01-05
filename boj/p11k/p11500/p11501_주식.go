package p11500

import (
	"fmt"
)

func Solve11501(reader Reader, writer Writer) {
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int64
		_, _ = fmt.Fscan(reader, &n)

		prices := make([]int64, n)
		for j := range prices {
			_, _ = fmt.Fscan(reader, &prices[j])
		}

		_, _ = fmt.Fprintln(writer, solve11501(prices))
	}
}

func solve11501(prices []int64) int64 {
	max := prices[len(prices)-1]
	res := int64(0)
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		} else {
			res += max - prices[i]
		}
	}
	return res
}
