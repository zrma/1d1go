package p7500

import (
	"fmt"
)

func Solve7579(reader Reader, writer Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)

	type appDatum struct {
		mem  int
		cost int
	}

	appData := make([]appDatum, n)
	for i := range appData {
		_, _ = fmt.Fscan(reader, &appData[i].mem)
	}
	for i := range appData {
		_, _ = fmt.Fscan(reader, &appData[i].cost)
	}

	const (
		maxCost   = 100
		maxCount  = 100
		costLimit = maxCost * maxCount
	)

	dp := make([]int, costLimit+1)
	for _, app := range appData {
		for cost := costLimit; cost >= app.cost; cost-- {
			v := dp[cost-app.cost] + app.mem
			if dp[cost] < v {
				dp[cost] = v
			}
		}
	}
	for cost, freeMem := range dp {
		if freeMem >= m {
			_, _ = fmt.Fprint(writer, cost)
			return
		}
	}
}
