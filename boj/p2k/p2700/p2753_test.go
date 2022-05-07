package p2700

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2753(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2753")

	for _, tt := range []struct {
		year int
		want int
	}{
		{2000, 1},
		{1999, 0},
		{2100, 0},
		{2104, 1},
		{2200, 0},
		{2400, 1},
	} {
		t.Run(fmt.Sprintf("%d", tt.year), func(t *testing.T) {
			got := Solve2753(tt.year)
			assert.Equal(t, tt.want, got)
		})
	}
}
