package p2400_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2400"
)

func TestSolve2480(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2480")

	for i, tt := range []struct {
		arr  [3]int
		want int
	}{
		{[3]int{3, 3, 6}, 1300},
		{[3]int{6, 3, 3}, 1300},
		{[3]int{2, 2, 2}, 12000},
		{[3]int{6, 2, 5}, 600},
		{[3]int{2, 6, 5}, 600},
		{[3]int{2, 5, 6}, 600},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := p2400.Solve2480(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
