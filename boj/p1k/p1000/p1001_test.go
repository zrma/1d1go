package p1000

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1001(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1001")

	for i, tt := range []struct {
		a, b int
		want int
	}{
		{3, 2, 3 - 2},
		{3, 2, 1},
		{1, 1, 1 - 1},
		{1, 1, 0},
		{9, 1, 9 - 1},
		{9, 1, 8},
		{1, 9, 1 - 9},
		{1, 9, -8},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve1001(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
