package p15500_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p15k/p15500"
)

func TestSolve15596(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/15596")

	for i, tt := range []struct {
		arr  []int
		want int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			15,
		},
		{
			[]int{5, 4, 3, 2},
			14,
		},
		{
			[]int{3, 6, 9},
			18,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := p15500.Solve15596(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
