package p2700

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2742(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2742")

	for _, tt := range []struct {
		n    int
		want []int
	}{
		{1, []int{1}},
		{2, []int{2, 1}},
		{3, []int{3, 2, 1}},
		{4, []int{4, 3, 2, 1}},
		{5, []int{5, 4, 3, 2, 1}},
		{6, []int{6, 5, 4, 3, 2, 1}},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := Solve2742(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
