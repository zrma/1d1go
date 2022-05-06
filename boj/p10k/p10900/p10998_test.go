package p10900

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve10998(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/10998")

	for i, tt := range []struct {
		a, b int
		want int
	}{
		{1, 1, 1 * 1},
		{1, 1, 1},
		{1, 2, 1 * 2},
		{1, 2, 2},
		{1, 9, 1 * 9},
		{1, 9, 9},
		{9, 1, 9 * 1},
		{9, 1, 9},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve10998(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
