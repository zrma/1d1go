package p2400

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2437(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2437")

	for i, tt := range []struct {
		arr  []int
		want int
	}{
		{[]int{3, 1, 6, 2, 7, 30, 1}, 21},
		{[]int{1, 1, 2, 3, 6, 7, 30}, 21},
		{[]int{1}, 2},
		{[]int{2}, 1},
		{[]int{3}, 1},
		{[]int{1, 1}, 3},
		{[]int{1, 2}, 4},
		{[]int{1, 3}, 2},
		{[]int{1, 4}, 2},
		{[]int{2, 4}, 1},
		{[]int{1, 2, 4}, 8},
		{[]int{1, 2, 5}, 4},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve2437(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
