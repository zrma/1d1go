package p1100

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve1110(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/1110")

	for _, tt := range []struct {
		n    int
		want int
	}{
		{26, 4},
		{55, 3},
		{1, 60},
		{0, 1},
		{71, 12},
	} {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			got := Solve1110(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
