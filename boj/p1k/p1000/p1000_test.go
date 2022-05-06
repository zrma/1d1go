package p1000

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1000(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1000")

	for i, tt := range []struct {
		a, b int
		want int
	}{
		{1, 2, 1 + 2},
		{1, 2, 3},
		{1, 9, 1 + 9},
		{1, 9, 10},
		{9, 1, 9 + 1},
		{9, 1, 10},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve1000(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
