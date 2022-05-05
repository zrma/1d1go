package p2500

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve2583(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2583")

	for i, tt := range []struct {
		row, col int
		rects    []Rect
		want     []int
	}{
		{
			row: 5, col: 7,
			rects: []Rect{
				{0, 2, 4, 4},
				{1, 1, 2, 5},
				{4, 0, 6, 2},
			},
			want: []int{1, 7, 13},
		},
		{
			row: 100, col: 100,
			want: []int{10000},
		},
		{
			row: 100, col: 100,
			rects: []Rect{
				{0, 0, 1, 1},
			},
			want: []int{9999},
		},
		{
			row: 100, col: 100,
			rects: []Rect{
				{1, 0, 2, 100},
			},
			want: []int{100, 9800},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Solve2583(tt.row, tt.col, tt.rects)
			assert.Equal(t, tt.want, got)
		})
	}
}
