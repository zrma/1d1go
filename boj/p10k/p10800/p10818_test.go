package p10800_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10818(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10818")

	for i, tt := range []struct {
		arr  []int
		want [2]int
	}{
		{
			[]int{20, 10, 35, 30, 7},
			[2]int{7, 35},
		},
		{
			[]int{1, 2, 3, 5, 4, 6, 10, 9, 8, 7},
			[2]int{1, 10},
		},
		{
			[]int{2, 3, 5, 4, 6, 10, 9, 8, 7, 1},
			[2]int{1, 10},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := p10800.Solve10818(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
