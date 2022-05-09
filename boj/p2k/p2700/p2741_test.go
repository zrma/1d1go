package p2700

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2741(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2741")

	for _, tt := range []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
		{4, []int{1, 2, 3, 4}},
		{5, []int{1, 2, 3, 4, 5}},
		{6, []int{1, 2, 3, 4, 5, 6}},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := Solve2741(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
