package p1400

import (
	"fmt"
	"io"
	"sort"
)

func Solve1450(reader io.Reader, writer io.Writer) {
	var n, c int
	_, _ = fmt.Fscan(reader, &n, &c)

	weights := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &weights[i])
	}

	// 0 ~ n/2 까지의 무게 조합
	result0 := knapsack(weights, 0, n/2, c, 0)

	// n/2 ~ n 까지의 무게 조합
	result1 := knapsack(weights, n/2, n, c, 0)

	sort.Ints(result0)
	sort.Ints(result1)

	res := 0
	lo, hi := 0, len(result1)-1

	for {
		if lo >= len(result0) || hi < 0 {
			break
		}

		if result0[lo]+result1[hi] > c {
			hi--
			continue
		}

		res += hi + 1
		lo++
	}

	_, _ = fmt.Fprint(writer, res)
}

func knapsack(weights []int, start, end, c, sum int) []int {
	if sum > c {
		return nil
	}

	if start == end {
		return []int{sum}
	}

	return append(
		knapsack(weights, start+1, end, c, sum+weights[start]), // 넣어봄.
		knapsack(weights, start+1, end, c, sum)...,             // 빼봄.
	)
}
