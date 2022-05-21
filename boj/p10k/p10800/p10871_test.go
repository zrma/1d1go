package p10800_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p10k/p10800"
)

func TestSolve10871(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10871")

	for i, tt := range []struct {
		x    int
		arr  []int
		want []int
	}{
		{
			5,
			[]int{1, 10, 4, 9, 2, 3, 8, 5, 7, 6},
			[]int{1, 4, 2, 3},
		},
		{
			4,
			[]int{1, 10, 4, 9, 2, 3, 8, 5, 7, 6},
			[]int{1, 2, 3},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := p10800.Solve10871(tt.x, tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
