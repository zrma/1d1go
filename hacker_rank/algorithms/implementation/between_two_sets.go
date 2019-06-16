package implementation

import (
	"sort"
)

func getTotalX(a []int32, b []int32) int32 {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	num1 := a[len(a)-1]
	num2 := b[0]

	cond1 := func(num int32, arr []int32) bool {
		for _, n := range arr {
			if num%n != 0 {
				return false
			}
		}
		return true
	}

	cond2 := func(num int32, arr []int32) bool {
		for _, n := range arr {
			if n%num != 0 {
				return false
			}
		}
		return true
	}

	var cnt int32
	for i := num1; i <= num2; i++ {
		if !cond1(i, a) || !cond2(i, b) {
			continue
		}
		cnt++
	}
	return cnt
}
