package p2500_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p2k/p2500"
)

func TestSolve2583(t *testing.T) {
	t.Log("https://www.acmicpc.net/problem/2583")

	for i, tt := range []struct {
		row, col int
		rects    []p2500.Rect
		want     []int
	}{
		{
			row: 5, col: 7,
			rects: []p2500.Rect{
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
			rects: []p2500.Rect{
				{0, 0, 1, 1},
			},
			want: []int{9999},
		},
		{
			row: 100, col: 100,
			rects: []p2500.Rect{
				{1, 0, 2, 100},
			},
			want: []int{100, 9800},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			io := newIOWithMock(t)
			io.EXPECT().
				Scan(gomock.Any()).
				Do(func(args ...any) {
					assert.Len(t, args, 3)
					row, ok := args[0].(*int)
					assert.True(t, ok)
					col, ok := args[1].(*int)
					assert.True(t, ok)
					k, ok := args[2].(*int)
					assert.True(t, ok)
					*row = tt.row
					*col = tt.col
					*k = len(tt.rects)
				}).
				Return(0, nil)
			for _, rect := range tt.rects {
				r := rect // capture
				io.EXPECT().
					Scan(gomock.Any()).
					Do(func(args ...any) {
						assert.Len(t, args, 4)
						for i, n := range args {
							ptr, ok := n.(*int)
							assert.True(t, ok)
							*ptr = r[i]
						}
					}).
					Return(0, nil)
			}

			io.EXPECT().
				Println(len(tt.want)).
				Return(0, nil)
			for _, want := range tt.want {
				io.EXPECT().
					Println(want).
					Return(0, nil)
			}
			p2500.Solve2583(io)
		})
	}
}
