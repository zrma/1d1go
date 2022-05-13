package p1000_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"1d1go/boj/p1k/p1000"
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
			io := newIOWithMock(t)
			io.EXPECT().
				Scan(gomock.Any()).
				Do(func(args ...any) {
					assert.Len(t, args, 2)
					pA, ok := args[0].(*int)
					assert.True(t, ok)
					pB, ok := args[1].(*int)
					assert.True(t, ok)
					*pA = tt.a
					*pB = tt.b
				}).
				Return(0, nil)
			io.EXPECT().
				Println(tt.want).
				Return(0, nil)

			p1000.Solve1001(io)
		})
	}
}
