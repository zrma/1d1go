package p1000

import (
	"fmt"
	"sort"
)

func Solve1082(reader Reader, writer Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)

	arr := make([]int, n)
	type price struct {
		num  int
		cost int
	}
	prices := make([]price, n)
	for i := range prices {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		arr[i] = v
		prices[i] = price{i, v}
	}

	if n == 1 {
		_, _ = fmt.Fprint(writer, 0)
		return
	}

	sort.Slice(prices, func(i, j int) bool {
		return prices[i].cost < prices[j].cost
	})

	var m int
	_, _ = fmt.Fscan(reader, &m)

	res := make([]int, 0)
	cost := 0
	cheapest := prices[0]
	if cheapest.num == 0 {
		second := prices[1]
		if second.cost > m {
			_, _ = fmt.Fprint(writer, 0)
			return
		}

		res = append(res, second.num)
		cost += second.cost
	}

	for cost+cheapest.cost <= m {
		res = append(res, cheapest.num)
		cost += cheapest.cost
	}

	for i, v := range res {
		for j := n - 1; j >= 0; j-- {
			// 해당 슬롯의 값을 큰 수부터 바꿔 넣기 시도해본다.
			newCost := cost - arr[v] + arr[j]
			if newCost <= m {
				cost = newCost
				res[i] = j
				break
			}
		}
	}

	for _, v := range res {
		_, _ = fmt.Fprint(writer, v)
	}
}
